/*
Fredy Eleazar Cardona Banegas 20201002906
Algortimos y estruturas de datos (IS310-1900)
*/
package main

import "strconv"
import "fmt"

//Crando la Cola
type ColaPrioridad struct {
     cola*  this.Celda
}

//Propiedades de la cola
type Celda struct {
     var elemento  int
     var prioridad  int
     sig*  this.Celda
}
func Celda_1() * Celda {
    var this * Celda =  &Celda{};
    return this;
}
//Validando lo elementos de la cola
func ColaPrioridad_1() * ColaPrioridad {
    var this * ColaPrioridad =  &ColaPrioridad{};
    this.cola = this.Celda_1();
    this.cola.sig = nil;
    return this;
}
func (this *ColaPrioridad) vacia()bool {
    return (this.cola.sig == nil);
}
func (this *ColaPrioridad) primero()int {
    if (this.vacia()) {
    panic("Exception");
    }
    var  aux*  this.Celda;
    aux = this.cola;
    this.cola = this.cola.sig;
    return aux.sig.elemento;
}

//Funciones para el desencolado de la cola 
func (this *ColaPrioridad) primero_prioridad()int {
    if (this.vacia()) {
    panic("Exception");
    }
    return this.cola.sig.prioridad;
}
func (this *ColaPrioridad) inserta( elemento  int,  prioridad  int) {
    var  p*  this.Celda;
    var  q*  this.Celda;
    var  encontrado  bool = false;
    p = this.cola;
    for((p.sig != nil) && (!encontrado)) {
        if (p.sig.prioridad > prioridad) {
        encontrado = true;
        } else {
            p = p.sig;
        }
    }
    q = p.sig;
    p.sig = this.Celda_1();
    p = p.sig;
    p.elemento = elemento;
    p.prioridad = prioridad;
    p.sig = q;
}
func (this *ColaPrioridad) suprime() {
    if (this.vacia()) {
    panic("Exception");
    }
    this.cola = this.cola.sig;
}
//Imprimir datos de la cola
func (this *ColaPrioridad) toString() string {
    return "ColaPrioridad [cola=" + strconv.Itoa(this.cola) + "]";
}