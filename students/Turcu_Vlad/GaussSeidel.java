import java.util.Scanner;
public class GaussSeidel {
   
	static Scanner sc = new Scanner(System.in);
	
	public static double norma(double[] x, double[] y, int n) {
		double max = 0;
		for (int i = 0; i < n; i++) {
			if(max<Math.abs(x[i] - y[i]))max=Math.abs(x[i] - y[i]);
		}
		return max;
	}
	public static void main(String[] args) {

		int n;
		System.out.println("n= ");
		n = sc.nextInt();
		double a[][] = new double[n][n];
		int k = 1, kmax = 500;
		double y[] = new double[n];
		double x[] = new double[n];
		for (int i = 0; i < n; i++)
			x[i] = 0;
		double b[] = new double[n];
		double ep = 0.001;
		double nor;

		do {

			for (int i = 0; i < n; i++) {
				double c = 0;
				for (int j = 0; j < n; j++) {
					if (j != i) {
						if (i < j)
							c = c + a[i][j] * y[j];
						else
							c = c + a[i][j] * x[i];
					}
				}
				y[i] = (1 / a[i][i]) * (b[i] - c);
			}
			k++;

			nor = norma(x, y, n);
			for (int i = 0; i < n; i++) {
				x[i] = y[i];
			}

		} while ((nor >= ep) && (k < kmax));
		
        for(int i=0 ;i<n; i++)
        	System.out.println(x[i]);
	}

}


