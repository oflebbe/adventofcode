package de.oflebbe;

class Main {
    public static void main( String[] argv) {

        System.out.println( new Circle( 9,25).timestep()   );
        System.out.println( new Circle( 17,1104).timestep()   );
        System.out.println( new Circle( 21,6111).timestep()   );
        System.out.println( new Circle( 30,5807).timestep()   );
        long start = System.currentTimeMillis();
        System.out.println( new Circle2( 459,71790).timestep()   );
        long stop = System.currentTimeMillis();
        System.out.printf("Timing %d\n", (stop-start));
    }


}