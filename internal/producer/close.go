package producer

// Close closes the connection to Kafka.
func (p *Producer) Close() error {
	return p.Producer.Close()
}
