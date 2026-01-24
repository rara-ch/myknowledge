package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rara-ch/myknowledge.git/internal/database"
)

func addTopicHandler(s *state, cmd inputCommand) error {
	if len(cmd.args) < 1 {
		return errors.New("not enough arguments to create a new topic")
	}

	topicName := cmd.args[0]

	if topic, err := s.db.GetTopicByName(context.Background(), topicName); err != sql.ErrNoRows {
		if err == nil {
			return fmt.Errorf("'%s' already exists as a topic", topic.Name)
		} else {
			return err
		}
	}

	topicDescription := sql.NullString{
		String: "",
		Valid:  false,
	}

	if len(cmd.args) >= 2 {
		topicDescription = sql.NullString{String: cmd.args[1], Valid: true}
	}

	topic, err := s.db.CreateTopic(context.Background(), database.CreateTopicParams{
		ID:          uuid.New(),
		Name:        topicName,
		Description: topicDescription,
	})
	if err != nil {
		return err
	}

	fmt.Printf("New topic '%s' was created\n", topic.Name)
	return nil
}
