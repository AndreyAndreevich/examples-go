// Code generated by mockery v2.30.16. DO NOT EDIT.

package rabbitmq

import (
	context "context"

	amqp091 "github.com/rabbitmq/amqp091-go"

	mock "github.com/stretchr/testify/mock"
)

// ChannelMock is an autogenerated mock type for the Channel type
type ChannelMock struct {
	mock.Mock
}

// Consume provides a mock function with given fields: queue, consumer, autoAck, exclusive, noLocal, noWait, args
func (_m *ChannelMock) Consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp091.Table) (<-chan amqp091.Delivery, error) {
	ret := _m.Called(queue, consumer, autoAck, exclusive, noLocal, noWait, args)

	var r0 <-chan amqp091.Delivery
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, bool, bool, amqp091.Table) (<-chan amqp091.Delivery, error)); ok {
		return rf(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
	}
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, bool, bool, amqp091.Table) <-chan amqp091.Delivery); ok {
		r0 = rf(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan amqp091.Delivery)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, bool, bool, bool, bool, amqp091.Table) error); ok {
		r1 = rf(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeDeclare provides a mock function with given fields: name, kind, durable, autoDelete, internal, noWait, args
func (_m *ChannelMock) ExchangeDeclare(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args amqp091.Table) error {
	ret := _m.Called(name, kind, durable, autoDelete, internal, noWait, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, bool, bool, amqp091.Table) error); ok {
		r0 = rf(name, kind, durable, autoDelete, internal, noWait, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishWithContext provides a mock function with given fields: ctx, exchange, key, mandatory, immediate, msg
func (_m *ChannelMock) PublishWithContext(ctx context.Context, exchange string, key string, mandatory bool, immediate bool, msg amqp091.Publishing) error {
	ret := _m.Called(ctx, exchange, key, mandatory, immediate, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool, bool, amqp091.Publishing) error); ok {
		r0 = rf(ctx, exchange, key, mandatory, immediate, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueueBind provides a mock function with given fields: name, key, exchange, noWait, args
func (_m *ChannelMock) QueueBind(name string, key string, exchange string, noWait bool, args amqp091.Table) error {
	ret := _m.Called(name, key, exchange, noWait, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, bool, amqp091.Table) error); ok {
		r0 = rf(name, key, exchange, noWait, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueueDeclare provides a mock function with given fields: name, durable, autoDelete, exclusive, noWait, args
func (_m *ChannelMock) QueueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp091.Table) (amqp091.Queue, error) {
	ret := _m.Called(name, durable, autoDelete, exclusive, noWait, args)

	var r0 amqp091.Queue
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool, bool, bool, bool, amqp091.Table) (amqp091.Queue, error)); ok {
		return rf(name, durable, autoDelete, exclusive, noWait, args)
	}
	if rf, ok := ret.Get(0).(func(string, bool, bool, bool, bool, amqp091.Table) amqp091.Queue); ok {
		r0 = rf(name, durable, autoDelete, exclusive, noWait, args)
	} else {
		r0 = ret.Get(0).(amqp091.Queue)
	}

	if rf, ok := ret.Get(1).(func(string, bool, bool, bool, bool, amqp091.Table) error); ok {
		r1 = rf(name, durable, autoDelete, exclusive, noWait, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewChannelMock creates a new instance of ChannelMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChannelMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChannelMock {
	mock := &ChannelMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
