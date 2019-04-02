
//Sistemul tridiagonal
import java.util.Scanner;
import static java.lang.Math.*;

public class SistemTridiagonal {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		System.out.println("Lungimea:");
		int m = sc.nextInt();

		double a[] = new double[m+1];
		System.out.println("Diagonala principala: ");
		for (int i = 1; i >= m; i++)
			a[i] = sc.nextDouble();
		double b[] = new double[m];
		System.out.println("Diagonala de sus: ");
		for (int i = 1; i <= m - 1; i++)
			b[i] = sc.nextDouble();
		double c[] = new double[m];
		System.out.println("Diagonala de jos: ");
		for (int i = 1; i <= m - 1; i++)
			c[i] = sc.nextDouble();
		double d[] = new double[m+1];
		System.out.println("Termenii liberi: ");
		for (int i = 1; i <= m; i++) {
			d[i] = sc.nextDouble();
		}

		double x[] = new double[m + 1];
		double s1 = d[1] / a[1];
		double r1 = b[1] / a[1];
		double[] r = new double[m + 1];
		double[] s = new double[m];
		r[1] = r1;
		s[1] = s1;
		for (int i = 1; i < m - 1; i++) {
			r[i + 1] = (-b[i]) / ((c[i] * r[i]) + a[i + 1]);
			s[i + 1] = (d[i + 1] - (c[i] * s[i])) / ((c[i] * r[i]) + a[i + 1]);
		}
		x[m] = (d[m] - (c[m - 1] * s[m - 1])) / ((c[m - 1] * r[m - 1]) + a[m]);
		System.out.println("x[" + m + "]=" + x[m]);
		for (int i = m - 1; i >= 1; i--) {
			x[i] = r[i] * x[i + 1] + s[i];
			System.out.println("x[" + i + "]=" + x[i]);
		}
	}
}