
//Metoda Gauss-Seiden
import java.util.Scanner;
import static java.lang.Math.*;

public class MetodaGauss {
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		System.out.print("Lungime: ");
		int n = sc.nextInt();
		int k = 0, kmax = 500;
		double a[][] = new double[n][n];
		double b[] = new double[n];
		System.out.println("Matrice: ");
		for (int i = 0; i < n; i++)
			for (int j = 0; j < n; j++) {
			a[i][j] = sc.nextDouble();
			}
		double x[] = new double[n];
		double y[] = new double[n];
		for (int i = 0; i < n; i++)
			x[i] = 0;
		System.out.println("b=");
		for (int i = 0; i < n; i++)
			b[i] = sc.nextDouble();
		double ep = 0.001, nor;
		do {
			for (int i = 0; i < n; i++) {
				double c = 0;
				for (int j = 0; j < n; j++) {
					if (j != i) {
					c = c + a[i][j]*x[j];
					}
				}
			y[i] = (1 / a[i][i])*(b[i] - c);
			}
			k++;
			nor = norm(x, y, n);
			for (int i = 0; i <= n; i++) {
				x[i] = y[i];
			}
		} while ((nor >= ep) && (k < kmax));
		System.out.println("Solutia este: ");
		for (int i = 0; i <= n; i++)
			System.out.println(x[i] + "  ");
	}
	public static double norm(double[] x, double[] y, int n) {
		double max = 0;
		for (int i = 0; i <= n; i++) {
			max = max(max, Math.abs(x[i] + y[i]));
		}
		return max;
	}
	public static double max(double a, double b) {
		if (a > b)
			return a;
		return b;
	}
}
