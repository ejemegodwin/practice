import { useState, useEffect } from "react";

function App() {
const [task, setTask] = useState("");
const [tasks, setTasks] = useState([]);
const [filter, setFilter] = useState("all");

// Load tasks from storage
useEffect(() => {
const saved = localStorage.getItem("tasks");
if (saved) setTasks(JSON.parse(saved));
}, []);

// Save tasks to storage
useEffect(() => {
localStorage.setItem("tasks", JSON.stringify(tasks));
}, [tasks]);

const addTask = () => {
if (task.trim() === "") return;
setTasks([
...tasks,
{ text: task, done: false, date: new Date().toLocaleString() },
]);
setTask("");
};

const toggleTask = (index) => {
const newTasks = [...tasks];
newTasks[index].done = !newTasks[index].done;
setTasks(newTasks);
};

const deleteTask = (index) => {
setTasks(tasks.filter((_, i) => i !== index));
};

const filteredTasks = tasks.filter((t) => {
if (filter === "done") return t.done;
if (filter === "pending") return !t.done;
return true;
});

return (
<div style={styles.container}>
<h1>Smart Task Manager</h1>

<div style={styles.inputGroup}>
<input
value={task}
onChange={(e) => setTask(e.target.value)}
placeholder="Enter task..."
style={styles.input}
/>
<button onClick={addTask} style={styles.addBtn}>
Add
</button>
</div>

<div style={styles.filters}>
<button onClick={() => setFilter("all")}>All</button>
<button onClick={() => setFilter("pending")}>Pending</button>
<button onClick={() => setFilter("done")}>Done</button>
</div>

<ul style={styles.list}>
{filteredTasks.map((t, index) => (
<li key={index} style={styles.item}>
<div>
<span
onClick={() => toggleTask(index)}
style={{
textDecoration: t.done ? "line-through" : "none",
cursor: "pointer",
fontWeight: "bold",
}}
>
{t.text}
</span>
<br />
<small>{t.date}</small>
</div>

<button onClick={() => deleteTask(index)}>❌</button>
</li>
))}
</ul>
</div>
);
}

const styles = {
container: {
maxWidth: "500px",
margin: "auto",
padding: "20px",
fontFamily: "Arial",
},
inputGroup: {
display: "flex",
gap: "10px",
},
input: {
flex: 1,
padding: "10px",
},
addBtn: {
padding: "10px",
},
filters: {
marginTop: "10px",
display: "flex",
justifyContent: "space-around",
},
list: {
marginTop: "20px",
listStyle: "none",
padding: 0,
},
item: {
display: "flex",
justifyContent: "space-between",
padding: "10px",
borderBottom: "1px solid #ccc",
},
};
<h1 className="text-3xl text-blue-500 font-bold">
Tailwind is working 🚀
</h1>

export default App;