package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ================================================================
// Lesson 32 — Kafka 消费者入门（Sarama 模式）
// ================================================================
//
// Kafka 是什么？
//   Kafka 是一个消息队列，像一条"消息管道"：
//   生产者（Producer）→ [Topic] → 消费者（Consumer）
//
// 核心概念：
//   - Topic（主题）：消息的分类，如 "marker_tool"、"prod_mp3"
//   - Partition（分区）：一个 Topic 可以拆成多个分区，并行消费
//   - Consumer Group（消费者组）：同一个组内的消费者分摊消费一个 Topic
//   - Offset（偏移量）：记录每个消费者读到哪了，重启后可以接着读
//
// 对照 issue-consumer 中的 Sarama 用法：
//   consumer_group.go  → NewKafkaConsumer + ConsumerLauncher
//   mark_tool.go       → MarkToolHandler 实现 sarama.ConsumerGroupHandler
//
// ================================================================

// ============================================================
// 模拟 Kafka 消息和消费者组
// ============================================================

// KafkaMessage 模拟一条 Kafka 消息。
type KafkaMessage struct {
	Topic     string
	Key       string
	Value     string
	Partition int
	Offset    int64
}

// MessageHandler 模拟 sarama.ConsumerGroupHandler 接口。
// 和 issue-consumer 的 MarkToolHandler 一样，需要实现三个方法。
type MessageHandler interface {
	Setup(claim PartitionClaim) error
	Cleanup(claim PartitionClaim) error
	ConsumeClaim(ctx context.Context, claim PartitionClaim) error
}

// PartitionClaim 模拟 Sarama 的 ConsumerGroupClaim。
// 核心是 Messages() channel：从 Kafka 来的消息通过这个 channel 传递。
type PartitionClaim struct {
	Topic     string
	Partition int
	messages  chan KafkaMessage
}

// Messages 返回消息 channel。
func (pc PartitionClaim) Messages() <-chan KafkaMessage {
	return pc.messages
}

// ============================================================
// 模拟的消费者组（相当于 Sarama 的 ConsumerGroup.Consume）
// ============================================================

// SimulatedConsumerGroup 模拟 Sarama 的消费者组。
// 它会创建多个 goroutine 模拟从多个分区读取消息，并回调你的 handler。
type SimulatedConsumerGroup struct {
	messages []KafkaMessage // 模拟 Kafka 中已有的消息
}

// Consume 模拟 Sarama 的 Consume 方法。
// 它把消息分配到各个"分区"，然后回调 handler 的 Setup → ConsumeClaim → Cleanup。
func (cg *SimulatedConsumerGroup) Consume(ctx context.Context, handler MessageHandler) error {
	// 模拟有 2 个分区
	partitionCount := 2

	// 把消息按分区分配
	partitions := make([]PartitionClaim, partitionCount)
	for i := 0; i < partitionCount; i++ {
		partitions[i] = PartitionClaim{
			Topic:     "test-topic",
			Partition: i,
			messages:  make(chan KafkaMessage, 10),
		}
	}

	// 把消息放到对应分区的 channel 里
	for _, msg := range cg.messages {
		p := msg.Partition
		partitions[p].messages <- msg
	}

	// 关闭所有 channel，表示"没有新消息了"
	for i := range partitions {
		close(partitions[i].messages)
	}

	// 对每个分区，回调 handler
	var wg sync.WaitGroup
	for _, claim := range partitions {
		wg.Add(1)
		go func(c PartitionClaim) {
			defer wg.Done()

			// 第 1 步：调 Setup
			_ = handler.Setup(c)

			// 第 2 步：调 ConsumeClaim（你的业务代码！）
			_ = handler.ConsumeClaim(ctx, c)

			// 第 3 步：调 Cleanup
			_ = handler.Cleanup(c)
		}(claim)
	}

	wg.Wait()
	return nil
}

// ============================================================
// 你的业务 Handler（模拟 issue-consumer 的 MarkToolHandler）
// ============================================================

// OrderHandler 处理订单消息。
// 实现 MessageHandler 接口：Setup / Cleanup / ConsumeClaim。
type OrderHandler struct {
	processedCount int
	mu             sync.Mutex
}

func (h *OrderHandler) Setup(claim PartitionClaim) error {
	fmt.Printf("[Setup]   分区 %d 开始消费\n", claim.Partition)
	return nil
}

func (h *OrderHandler) Cleanup(claim PartitionClaim) error {
	fmt.Printf("[Cleanup] 分区 %d 消费完成\n", claim.Partition)
	return nil
}

func (h *OrderHandler) ConsumeClaim(ctx context.Context, claim PartitionClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				// channel 关闭了，没有更多消息
				fmt.Printf("[ConsumeClaim] 分区 %d 消息读完\n", claim.Partition)
				return nil
			}
			// 处理消息
			h.mu.Lock()
			h.processedCount++
			h.mu.Unlock()
			fmt.Printf("[ConsumeClaim] 分区 %d 收到消息: offset=%d, key=%s\n",
				claim.Partition, msg.Offset, msg.Key)
			time.Sleep(50 * time.Millisecond) // 模拟处理耗时

		case <-ctx.Done():
			fmt.Printf("[ConsumeClaim] 分区 %d 收到取消信号，停止消费\n", claim.Partition)
			return ctx.Err()
		}
	}
}

// ============================================================
// main
// ============================================================

func main() {
	fmt.Println("========== Lesson 32: Kafka 消费者入门（Sarama 模式）==========")
	fmt.Println()

	// ============ 1. 模拟 Kafka 消息 ============
	fmt.Println("--- 1. 模拟 Kafka 中的消息 ---")
	messages := []KafkaMessage{
		{Topic: "test-topic", Key: "order:001", Value: `{"order_id":1,"amount":100}`, Partition: 0, Offset: 0},
		{Topic: "test-topic", Key: "order:002", Value: `{"order_id":2,"amount":200}`, Partition: 1, Offset: 0},
		{Topic: "test-topic", Key: "order:003", Value: `{"order_id":3,"amount":300}`, Partition: 0, Offset: 1},
		{Topic: "test-topic", Key: "order:004", Value: `{"order_id":4,"amount":400}`, Partition: 1, Offset: 1},
		{Topic: "test-topic", Key: "order:005", Value: `{"order_id":5,"amount":500}`, Partition: 0, Offset: 2},
	}
	fmt.Printf("共 %d 条消息，分布在 2 个分区\n", len(messages))
	fmt.Println()

	// ============ 2. 创建消费者组和 Handler ============
	fmt.Println("--- 2. 消费消息 ---")
	cg := &SimulatedConsumerGroup{messages: messages}
	handler := &OrderHandler{}

	ctx := context.Background()
	_ = cg.Consume(ctx, handler)

	fmt.Println()
	fmt.Printf("共处理了 %d 条消息\n", handler.processedCount)
	fmt.Println()

	// ============ 3. 对照 issue-consumer ============
	fmt.Println("--- 3. 对照 issue-consumer 真实代码 ---")
	fmt.Println()
	fmt.Println("Sarama 的 ConsumerGroupHandler 接口（你实现的部分）：")
	fmt.Println("  Setup(ConsumerGroupSession) error")
	fmt.Println("  Cleanup(ConsumerGroupSession) error")
	fmt.Println("  ConsumeClaim(ConsumerGroupSession, ConsumerGroupClaim) error")
	fmt.Println()
	fmt.Println("对应 issue-consumer 中的 MarkToolHandler：")
	fmt.Println("  biz/mark_tool.go:31 → Setup")
	fmt.Println("  biz/mark_tool.go:35 → Cleanup")
	fmt.Println("  biz/mark_tool.go:39 → ConsumeClaim（核心！逐条处理消息）")
	fmt.Println()
	fmt.Println("consumer_group.go 中的 Run 方法：")
	fmt.Println("  cl.client.Consume(ctx, topics, handler)")
	fmt.Println("  → Sarama 内部：分配分区 → 回调 handler.Setup()")
	fmt.Println("  → 对每个分区：回调 handler.ConsumeClaim()")
	fmt.Println("  → 分区消费完：回调 handler.Cleanup()")
	fmt.Println()

	// ============ 4. 总结 ============
	fmt.Println("========== 总结 ==========")
	fmt.Println()
	fmt.Println("Kafka 消费者模式的核心：")
	fmt.Println("  1. 你实现了一个接口（ConsumerGroupHandler）")
	fmt.Println("  2. 接口有 3 个方法：Setup / Cleanup / ConsumeClaim")
	fmt.Println("  3. Sarama 库负责：连接 Kafka、分配分区、读取消息")
	fmt.Println("  4. Sarama 回调你的方法：Setup → ConsumeClaim → Cleanup")
	fmt.Println()
	fmt.Println("消息流转：")
	fmt.Println("  Kafka Brokers → Sarama 库 → claim.Messages() channel → 你的 ConsumeClaim")
	fmt.Println()
	fmt.Println("issue-consumer 中的完整链路：")
	fmt.Println("  main.go")
	fmt.Println("    → NewKafkaConsumer(config)    // 创建 Kafka 连接")
	fmt.Println("    → NewMarkToolHandler(config)   // 创建你的业务 handler")
	fmt.Println("    → NewMConsumerLauncher(cg, handler, topics)")
	fmt.Println("    → Run(wg, ctx)                // 启动消费")
	fmt.Println("      → client.Consume(ctx, topics, handler)")
	fmt.Println("        → handler.Setup()")
	fmt.Println("        → handler.ConsumeClaim()  // 你的业务代码！")
	fmt.Println("        → handler.Cleanup()")
}