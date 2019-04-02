import java.util.Scanner;
public class Jacobi {

	public static double norma(double[] x, double[] y, int n) {
		double max = 0;
		for (int i = 0; i < n; i++) {
			if(max<Math.abs(x[i] - y[i]))max=Math.abs(x[i] - y[i]);
		}
		return max;
	}
    static Scanner sc = new Scanner(System.in);
	public static void main(String[] args) {
		int n;
	
		System.out.println("n= ");
		n = sc.nextInt();
		double a[][] = new double[n][n];double x[] = new double[n];double y[] = new double[n];double b[] = new double[n];
		int k = 1, kmax = 500;
		System.out.println("a[][]=");
		for (int i = 0; i < n; i++)
		for (int j = 0; j < n; j++) {
		a[i][j] = sc.nextDouble();
			}
		for (int i = 0; i < n; i++)x[i] = 0;
		
		System.out.println("b[]=");
		for (int i = 0; i < n; i++)
			b[i] = sc.nextDouble();
		double ep = 0.001;
		double nor;

		do {
			for (int i = 0; i < n; i++) {
				double c = 0;
				for (int j = 0; j < n; j++) {
					if (j != i) {
						c = c + a[i][j] * x[j];
					}}
				y[i] = (1 / a[i][i]) * (b[i] - c);
			}
			k++;

			nor = norma(x, y, n);
			for (int i = 0; i < n; i++) {
				x[i] = y[i];
			}

		} while ((nor >= ep) && (k < kmax));

		System.out.println("Rezultat=");
		for (int i = 0; i < n; i++)
			System.out.println(x[i] + " ");
	}

	
}
