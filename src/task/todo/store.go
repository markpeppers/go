package todo

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/boltdb/bolt"
)

type Store struct {
	db bolt.DB
}

func NewStore() (*Store, error) {
	homeDirName, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	taskPath := path.Join(homeDirName, ".task")
	_, err = os.Stat(taskPath)
	if os.IsNotExist(err) {
		fmt.Println("Creating the directory")
		err = os.Mkdir(taskPath, 0755)
	}
	newDb, err := bolt.Open(path.Join(taskPath, "tasks.db"), 0600, nil)
	if err != nil {
		return nil, err
	}
	return &Store{db: *newDb}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) CreateTask(desc string) error {
	t := Task{Description: desc}
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return err
		}
		id, _ := b.NextSequence()
		t.Id = int(id)
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}
		return b.Put(itob(t.Id), buf)
	})
}

// Returns non-finished tasks
func (s *Store) GetTasks() ([]Task, error) {
	tasks := make([]Task, 0)
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		if b == nil {
			return errors.New("ERROR: No tasks yet 🥺")
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := &Task{}
			err := json.Unmarshal(v, task)
			if err != nil {
				return err
			}
			if !task.Finished {
				tasks = append(tasks, *task)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s Store) GetFinishedTasks(since time.Time) ([]Task, error) {
	tasks := make([]Task, 0)
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		if b == nil {
			return errors.New("Error: No tasks yet")
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := &Task{}
			err := json.Unmarshal(v, task)
			if err != nil {
				return err
			}
			if task.Finished && task.FinishTime.After(since) {
				tasks = append(tasks, *task)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Store) DeleteTask(n int) (string, error) {
	tasks, err := s.GetTasks()
	if err != nil {
		return "", err
	}
	if n > len(tasks) {
		return "", errors.New(fmt.Sprintf("No such task %d", n))
	}
	taskId := tasks[n-1].Id
	removedTask := tasks[n-1].Description
	err = s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		if b == nil {
			return errors.New("ERROR: Can't remove task from empty list")
		}
		return b.Delete(itob(taskId))
	})
	return removedTask, err
}

func (s *Store) FinishTask(n int) (string, error) {
	tasks, err := s.GetTasks()
	if err != nil {
		return "", err
	}
	if n > len(tasks) {
		return "", errors.New(fmt.Sprintf("No such task %d", n))
	}
	task := tasks[n-1]
	err = s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		if b == nil {
			return errors.New("ERROR: Can't remove task from empty list")
		}

		task.Finished = true
		task.FinishTime = time.Now()
		buf, err := json.Marshal(task)
		if err != nil {
			return err
		}
		b.Delete(itob(task.Id))
		b.Put(itob(task.Id), buf)
		return nil
	})
	return task.Description, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
