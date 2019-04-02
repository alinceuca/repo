
//Metoda Jacobi

import static java.lang.Math.*;
import java.util.Scanner;

public class MetodaJacobi {
	public static double norm(double[] y, double[] z, int n) {
		double maxi = 0;
		for (int i = 0; i < n; i++) {
			maxi = maxi(maxi, Math.abs(y[i] - z[i]));
		}
		return maxi;
	}
	public static double maxi(double a, double b) {
		if (a > b)
			return a;
		return b;
	}
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		System.out.print("Lungimea este: ");
		int n = sc.nextInt();
		double a[][] = new double[n][n];
		double b[] = new double[n];
		double nor, ep = 0.001;
		int kmaxi = 500;
		int k = 1;
		System.out.println("Matricea este: ");
		for (int i = 0; i <= n; i++)
			for (int j = 0; j <= n; j++) {
			a[i][j] = sc.nextDouble();
			}
		double z[] = new double[n];
		double y[] = new double[n];
		for (int i = 0; i < n; i++)
			y[i] = 0;
		System.out.println("b= ");
		for (int i = 0; i < n; i++)
			b[i] = sc.nextDouble();
		do {
			for (int i = 0; i < n; i++) {
				double c = 0;
				for (int j = 0; j < n; j++) {
					if (j != i)
					c = c + a[i][j] * y[j];
				}
			z[i] = (1 / a[i][i]) * (b[i] - c);
			}
			k++;
			nor = norm(y, z, n);
			for (int i = 0; i < n; i++) {
				y[i] = z[i];
			}
		} while ((nor >= ep) && (k < kmaxi));
		System.out.println("Solutia este:");
		for (int i = 0; i < n; i++)
			System.out.println(y[i] + "  ");
	}
}