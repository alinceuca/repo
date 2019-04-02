import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.util.Scanner;

public class Jacobi {
//http://www.csc.kth.se/utbildning/kth/kurser/DN1213/numme06/utdelat/kap6.pdf
	private double matrixA[][];
	private double b[],x[],y[];
	private int n;
	private double eps;
	private long nr_max;
	
	
	private double calcNorm() {
		double max = Math.abs(x[0] - y[0]);
		for(int i=1;i<n;i++) {
			double calc = Math.abs(x[i] - y[i]);
			if(calc > max)
				max = calc;
		}
		return max;
	}
	
	private double calcSum(int row) {
		double sum = 0.0;
		
		for(int i = 0;i<n;i++)
			if(i != row)
				sum += matrixA[row][i] * x[i];
		
		return sum;
	}
	
	public Jacobi(String path,long nr_max,double eps) {
		try {
			this.eps = eps;
			this.nr_max = nr_max;
			Scanner sc = new Scanner(new FileReader(new File(path)));
			n = sc.nextInt();
			
			matrixA = new double[n][n];
			b = new double[n];
			x = new double[n];
			y = new double[n];
			double norm;
			
			for(int i =0 ;i<n;i++) 
				for(int j =0 ;j<n;j++) 
					matrixA[i][j] = Double.parseDouble(sc.next());
			for(int i =0 ;i<n;i++)
					b[i] = Double.parseDouble(sc.next());
			sc.close();
			
			
			int nr_it = 0;
			
			do {
				
				for(int i=0;i<n;i++) 
					y[i] = (1.00/matrixA[i][i]) * (b[i] - calcSum(i));
				
				norm  = calcNorm();
				
				
				for(int i=0;i<n;i++)
					x[i] = y[i];
				
			
				
				nr_it++;
				
			}while(norm >= eps && nr_it < nr_max);
			
			if(norm < eps)
				System.out.println("Precision ok");
			else
				System.out.println("Not at full precision");
			
			for(int i=0;i<n;i++)
				System.out.println(y[i]);
			
			
		} catch (FileNotFoundException e) {

			e.printStackTrace();
		}
		
	}
}
