package internal

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

type Port struct {
	Name     string `json:"name,omitempty"`
	StreamID string `json:"stream_id,omitempty"`
}

type ActionGraph struct {
	Actions []*Action `json:"actions,omitempty"`
	Outputs []*Port   `json:"outputs,omitempty"`
}

type Action struct {
	Name    string  `json:"name,omitempty"`
	Inputs  []*Port `json:"inputs,omitempty"`
	Outputs []*Port `json:"inputsomitempty"`
	// TODO: Add configs.
}

type Chunk struct {
	MIMEType string `json:"mime_type,omitempty"`
	Data     []byte `json:"data,omitempty"`
	// TODO: Add metadata.
}

type StreamFrame struct {
	StreamID  string `json:"stream_id,omitempty"`
	Data      *Chunk `json:"chunk,omitempty"`
	Continued bool   `json:"continued,omitempty"`
}

type executeActionsMsg struct {
	SessionID    string         `json:"session_id,omitempty"`
	StreamFrames []*StreamFrame `json:"stream_frames,omitempty"`
}

func NewClient(endpoint string, apiKey string) (*Client, error) {
	c, _, err := websocket.DefaultDialer.Dial(endpoint+"?key="+apiKey, nil)
	if err != nil {
		return nil, err
	}
	return &Client{conn: c}, nil
}

type Session struct {
	c         *Client
	sessionID string
}

func (c *Client) OpenSession(sessionID string) (*Session, error) {
	// if err := c.conn.WriteJSON(&startSessionRequest{
	// 	// ProposedID: proposedID,
	// }); err != nil {
	// 	return nil, err
	// }
	// var resp startSessionResponse
	// if err := c.conn.ReadJSON(&resp); err != nil {
	// 	return nil, err
	// }
	// TODO(jbd) Start session for real.
	return &Session{c: c, sessionID: sessionID}, nil
}

func (s *Session) ExecuteActions(actions []*Action, outputs []string) error {
	if err := s.c.conn.WriteJSON(&executeActionsMsg{
		SessionID: s.sessionID,
		StreamFrames: []*StreamFrame{
			{StreamID: "test", Data: &Chunk{MIMEType: "text/plain", Data: []byte("hello world")}},
		},
	}); err != nil {
		return err
	}
	panic("not yet")
}
