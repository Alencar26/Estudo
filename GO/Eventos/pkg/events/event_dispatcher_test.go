package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name     string
	Patyload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Patyload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.eventDispatcher = NewEventDispatcher()
	s.event = TestEvent{Name: "Event", Patyload: "TesteEvent"}
	s.event2 = TestEvent{Name: "Event2", Patyload: "TesteEvent2"}
	s.handler = TestEventHandler{ID: 1}
	s.handler2 = TestEventHandler{ID: 2}
	s.handler3 = TestEventHandler{ID: 3}
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := s.eventDispatcher.Register(s.event.Name, &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.Name, &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	assert.Equal(s.T(), &s.handler, s.eventDispatcher.handlers[s.event.GetName()][0])
	assert.Equal(s.T(), &s.handler2, s.eventDispatcher.handlers[s.event.GetName()][1])
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := s.eventDispatcher.Register(s.event.Name, &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.Name, &s.handler)
	s.Equal(ErrHandlerAlreadyRegister, err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	//Event1
	err := s.eventDispatcher.Register(s.event.Name, &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.Name, &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	//Event2
	err = s.eventDispatcher.Register(s.event2.Name, &s.handler3)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event2.GetName()]))

	s.eventDispatcher.Clear()
	s.Equal(0, len(s.eventDispatcher.handlers))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	//Event1
	err := s.eventDispatcher.Register(s.event.Name, &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.Name, &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	assert.True(s.T(), s.eventDispatcher.Has(s.event.GetName(), &s.handler))
	assert.True(s.T(), s.eventDispatcher.Has(s.event.GetName(), &s.handler2))
	assert.False(s.T(), s.eventDispatcher.Has(s.event.GetName(), &s.handler3))
}

type MockedEventHandler struct {
	mock.Mock
}

func (m *MockedEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockedEventHandler{}
	eh.On("Handle", &s.event)
	s.eventDispatcher.Register(s.event.GetName(), eh)
	s.eventDispatcher.Dispatch(&s.event)
	eh.AssertExpectations(s.T())
	eh.AssertNumberOfCalls(s.T(), "Handle", 1)
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	//Event1
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	//Event2'
	err = s.eventDispatcher.Register(s.event2.GetName(), &s.handler3)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event2.GetName()]))

	//REMOVE HANDLER1
	s.eventDispatcher.Remove(s.event.GetName(), &s.handler)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
	assert.Equal(s.T(), &s.handler2, s.eventDispatcher.handlers[s.event.GetName()][0])

	//REMOVE HANDLER2
	s.eventDispatcher.Remove(s.event.GetName(), &s.handler2)
	s.Equal(0, len(s.eventDispatcher.handlers[s.event.GetName()]))

	//REMOVE HANDLER3
	s.eventDispatcher.Remove(s.event2.GetName(), &s.handler3)
	s.Equal(0, len(s.eventDispatcher.handlers[s.event2.GetName()]))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
