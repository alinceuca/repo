import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.util.Scanner;

public class Main {

		
	
		public static void main(String args[]) {
				
		
			Tridiagonal tri = new Tridiagonal("tridiagonal.in");
			Tridiagonal trr = new Tridiagonal("tridiagonal.in");
			
			//tri.solveIterative();
			
			//trr.solveRecursive(1,false);
			//trr.showSolutions();
			
			Jacobi jcb = new Jacobi(		"jacobi.in",10000, 0.0001);
			GausSeitel gs = new GausSeitel("jacobi.in", 10000, 0.0001);
		}
}
