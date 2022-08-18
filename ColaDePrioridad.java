public class ColaPrioridad{
    class Celda {
        int elemento;
        int prioridad;
        Celda sig;
    }
    private Celda cola;
    public ColaPrioridad() {
        cola = new Celda();
        cola.sig = null;
    }
    public boolean vacia() {
        return (cola.sig==null);
    }
    public int primero() throws Exception {
        if (vacia()) throw new Exception();
        Celda aux;
        aux=cola;
        cola=cola.sig;
        return aux.sig.elemento;
    }
    public int primero_prioridad() throws Exception {
        if (vacia()) throw new Exception();
        return cola.sig.prioridad;
    }
    public void inserta(int elemento, int prioridad) {
        Celda p,q;
        boolean encontrado = false;
        p = cola;
        while((p.sig!=null)&&(!encontrado)) {
            if (p.sig.prioridad>prioridad)
                encontrado = true;
            else p = p.sig;
        }
        q = p.sig;
        p.sig = new Celda();
        p = p.sig;
        p.elemento = elemento;
        p.prioridad = prioridad;
        p.sig = q;
    }
        public void suprime() throws Exception {
        if (vacia()) throw new Exception();
        cola = cola.sig;
    }
        @Override
        public String toString() {
            return "ColaPrioridad [cola=" + cola + "]";
        }

    
}