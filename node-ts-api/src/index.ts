import express from 'express';
import taskRoutes from './routes/tasks';

const app = express();
const port = 3000;

app.use(express.json());
app.use('/tasks', taskRoutes);

app.listen(port, () => {
  console.log(`✅ Task Manager API running at http://localhost:${port}`);
});
