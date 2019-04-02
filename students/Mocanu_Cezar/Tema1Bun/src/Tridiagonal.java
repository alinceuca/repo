import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.text.DecimalFormat;
import java.util.Scanner;

public class Tridiagonal {

	private double a[],b[],c[],d[],x[],r[],s[];
	private int n;
	private DecimalFormat df;
	
	public Tridiagonal(String path){
		df = new DecimalFormat("0.00");
		
		try {
			readData(path);
		} catch (IOException e) {
			
			e.printStackTrace();
		}
		x = new double[n];
		r = new double[n];
		s = new double[n];
		
		r[0] = -1 * ( c[0] / a[0] );
		s[0] = d[0] / a[0];
	}
	
	public void showSolutions() {
		System.out.println();
		for(int i=0;i<n;i++) 
			System.out.println("x" + (i+1) + " : "+  df.format(x[i]));
	}

	public void readData(String path) throws IOException{
		Scanner sc = new Scanner(new FileReader(new File(path)));
		n = sc.nextInt();
		
		a = new double[n];
		b = new double[n];
		c = new double[n];
		d = new double[n];
		
		for(int i=0;i<n;i++) {
			a[i] = Double.parseDouble(sc.next());
		}	
		
		for(int i=0;i<n-1;i++) {
			b[i] = Double.parseDouble(sc.next());
		}
		
		for(int i=0;i<n-1;i++) {
			c[i] = Double.parseDouble(sc.next());
		}
		
		for(int i=0;i<n;i++) {
			d[i] = Double.parseDouble(sc.next());
		}
		
		sc.close();
	}
	
	public void solveIterative() {
		for(int i=1;i<n-1;i++) {
			r[i] = (-1 * c[i]) / (b[i] * r[i-1] + a[i]);
			s[i] = (d[i] - b[i] * s[i-1]) / (b[i] * r[i-1] + a[i]);
		}
		x[n-1] = (d[n-1] - b[n-2] * s[n-2])/(b[n-2] * r[n-2] + a[n-1]);
		for(int i=n-2;i>=0;i--) 
			x[i] = r[i] * x[i+1] + s[i];
		
		showSolutions();
	}
	
	
	public void solveRecursive(int i,boolean backward) {
		
		if(!backward) {
			if(i < n-1) {
				r[i] = (-1 * c[i]) / (b[i] * r[i-1] + a[i]);
				s[i] = (d[i] - b[i] * s[i-1]) / (b[i] * r[i-1] + a[i]);
				solveRecursive(i+1,false);
			}	
			else {
				x[n-1] = (d[n-1] - b[n-2] * s[n-2])/(b[n-2] * r[n-2] + a[n-1]);
				solveRecursive(i-1,true);
			}
		}
		else if(i >= 0 && backward) {
			x[i] = r[i] * x[i+1] + s[i];
			solveRecursive(i-1,true);
		}
		
		
	}
	
}
