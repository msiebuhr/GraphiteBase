package GraphiteBase

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func parseGraphiteLine(line string) (*Metric, error) {
	// Find newline-rune
	fields := strings.Fields(line)

	// A line must at least contain <metric.name> <timestamp> <value> <tag=value>+
	if len(fields) != 3 {
		return new(Metric), errors.New("Invalid line")
	}

	// Convert name
	name := string(fields[0])

	// Parse out value
	value, err := strconv.ParseFloat(string(fields[1]), 64)
	if err != nil {
		return new(Metric), err
	}

	// Parse out timestamp
	time, err := strconv.ParseInt(string(fields[2]), 10, 64)
	if err != nil {
		return new(Metric), err
	}

	// Return new metric point
	return NewMetric(name, time, value), nil
}

func GraphiteProtocolReader(conn io.ReadCloser, out chan *Metric) error {
	scanner := bufio.NewScanner(conn)
	defer conn.Close()
	defer close(out)

	// Parse lines and hand them to the back-end
	for scanner.Scan() {
		m, err := parseGraphiteLine(scanner.Text())
		if err != nil {
			return err
		}
		out <- m
	}

	// Catch scanner errors too
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
