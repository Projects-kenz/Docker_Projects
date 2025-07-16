import express from 'express';
import { Task } from '../models/Task';
import { v4 as uuidv4 } from 'uuid';

const router = express.Router();

let tasks: Task[] = [];

router.get('/', (_req, res) => {
  res.json(tasks);
});

router.post('/', (req, res) => {
  const { title } = req.body;

  if (!title) {
    return res.status(400).json({ error: 'Title is required' });
  }

  const newTask: Task = {
    id: uuidv4(),
    title,
    completed: false
  };

  tasks.push(newTask);
  res.status(201).json(newTask);
});

router.delete('/:id', (req, res) => {
  const { id } = req.params;
  const index = tasks.findIndex(task => task.id === id);

  if (index === -1) {
    return res.status(404).json({ error: 'Task not found' });
  }

  const deleted = tasks.splice(index, 1);
  res.json(deleted[0]);
});

export default router;
