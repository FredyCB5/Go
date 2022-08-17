/*
Fredy Eleazr Cardona Banegas 20201002906
Algortimos y estruturas de datos (IS310-1900)
*/

package main

import "strconv"
import "fmt"


//Propiedades de la cola
type Cola struct {
     MAXTAM  int = 10;
     frente  int
     ultimo  int
     cola    Object
}

//Creando cola
func Cola_1() * Cola {
    var this * Cola =  &Cola{};
    this.frente = 0;
    this.ultimo = -1;
    this.cola = [this.MAXTAM]Object;
    return this;
}

//Metodo para insertar en la cola elementos en la cola.(Encolando)
func (this *Cola) insertar( elemento*  Object) {
    if (!this.ColaLlena()) {
    this.cola[this.ultimo++] = elemento;
    } else {
        panic("Error: Cola LLena"); 
    }
}
//Metodo para eliminar elementos de la cola.(Desencolando)

func (this *Cola) eliminar()* Object {
    if (!this.ColaVacia()) {
        return this.cola[this.frente++];
    } else {
        panic("La cola esta vacia");
    }
}
//Metodo para eliminar la cola
func (this *Cola) Borrar() {
    this.frente = 0;
    this.ultimo = -1;
}
func (this *Cola) frente()* Object {
    if (!this.ColaVacia()) {
        return this.cola[this.frente];
    } else {
        panic("Cola vacia");
    }
}
func (this *Cola) ColaVacia()bool {
    return this.frente > this.ultimo;
}
func (this *Cola) ColaLlena()bool {
    return this.ultimo == this.MAXTAM - 1;
}

//Imprimir la cola
func (this *Cola) toString() string {
    return "Cola [cola=" + Arrays.toString(this.cola) + ", frente=" + strconv.Itoa(this.frente) + ", ultimo=" + strconv.Itoa(this.ultimo) + "]";
}
