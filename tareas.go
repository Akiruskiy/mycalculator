package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) addTask(tl *task) { // agregar un elemento a un slice
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) deleteTask(i int) { // eliminar un elemento a un slice utilizando append
	t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
}

func (t *taskList) printList() {
	for _, ta := range t.tasks {
		fmt.Println("Nombre: ", ta.name, "\nDescripcion:", ta.description, "\n")
	}
}

func (t *taskList) printCompletedList() {
	for _, ta := range t.tasks {
		if ta.completed {
			fmt.Println("Nombre: ", ta.name, "\nDescripcion:", ta.description, "\n")
		}
	}
}

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) markCompleted() { // con * pudimos agarrar y cambiar los valores desde la memoria interna de la maquina
	t.completed = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}

func main() {
	t1 := &task{ // es el valor sobre el puntero para cambiar las variables por dentro
		name:        "Completar mi curso de GO",
		description: "Es el segundo curso de GO que estoy haciendo, hay que completarlo",
		completed:   false,
	}
	t2 := &task{ // es el valor sobre el puntero para cambiar las variables por dentro
		name:        "Completar mi curso de Kubernetes",
		description: "Es el segundo curso de k8 que estoy haciendo, hay que completarlo",
		completed:   false,
	}
	t3 := &task{ // es el valor sobre el puntero para cambiar las variables por dentro
		name:        "Completar mi curso de Dockers",
		description: "Es el segundo curso de DOckers que estoy haciendo, hay que completarlo",
		completed:   false,
	}
	/*fmt.Printf("%+v\n", t)
	t.markCompleted()
	t.updateDescription("Finalizar mi curso de GO")
	t.updateName("Curso GO")
	fmt.Printf("%+v\n", t)*/

	list := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	list.addTask(t3)
	list.printList()
	list.tasks[0].markCompleted()
	fmt.Println("Completed tasks:")
	list.printCompletedList()

	mapTasks := make(map[string]*taskList)
	mapTasks["Nestor"] = list

	t4 := &task{ // es el valor sobre el puntero para cambiar las variables por dentro
		name:        "Completar mi curso de Dockers",
		description: "Es el segundo curso de DOckers que estoy haciendo, hay que completarlo",
		completed:   false,
	}
	t5 := &task{ // es el valor sobre el puntero para cambiar las variables por dentro
		name:        "Completar mi curso de Dockers",
		description: "Es el segundo curso de DOckers que estoy haciendo, hay que completarlo",
		completed:   false,
	}
	list2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapTasks["Ricardo"] = list2

	fmt.Println("Tareas de Nestor")
	mapTasks["Nestor"].printList()
	fmt.Println("Tareas de Ricardo")
	mapTasks["Ricardo"].printList()

	/*for i := 0; i < len(list.tasks); i++ {
		fmt.Println("Index:", i, "nombre", list.tasks[i].name)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break // Para romper al iteracion
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // Para volver al for y no hacer lo que esta abajo
		}
		fmt.Println(i)
	}*/
}
