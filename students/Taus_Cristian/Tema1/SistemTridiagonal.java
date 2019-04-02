import java.util.Scanner;
public class SistemTridiagonal {

	public static void main(String[] args) {
		
		// TODO Auto-generated method stub
		Scanner sc = new Scanner(System.in);
		System.out.println("Introdu lungimea:");
		int n = sc.nextInt();

		double a[] = new double[n + 1];
		double b[] = new double[n];
		double c[] = new double[n];
		double d[] = new double[n + 1];
		double x[] = new double[n + 1];

		System.out.println("Introdu diagonala principala: ");
		for (int i = 1; i <= n; i++)
			a[i] = sc.nextDouble();
		System.out.println("Introdu diagonala de sus: ");
		for (int i = 1; i <= n - 1; i++)
			b[i] = sc.nextDouble();
		System.out.println("Introdu diagonala de jos: ");
		for (int i = 1; i <= n - 1; i++)
			c[i] = sc.nextDouble();
		System.out.println("Introdu termenii liberi: ");
		for (int i = 1; i <= n; i++) {
			d[i] = sc.nextDouble();
		}

		double r1 = -b[1] / a[1];
		double s1 = d[1] / a[1];


			double[] r = new double[n + 1];
			r[1] = r1;
			double[] s = new double[n];
			s[1] = s1;
			for (int i = 1; i < n - 1; i++) {
				r[i + 1] = (-b[i]) / ((c[i] * r[i]) + a[i + 1]);
				s[i + 1] = (d[i + 1] - (c[i] * s[i])) / ((c[i] * r[i]) + a[i + 1]);
			}
			x[n] = (d[n] - (c[n - 1] * s[n - 1])) / ((c[n - 1] * r[n - 1]) + a[n]);
			System.out.println("x[" + n + "]=" + x[n]);
			for (int i = n - 1; i >= 1; i--) {
				x[i] = r[i] * x[i + 1] + s[i];
				System.out.println("x[" + i + "]=" + x[i]);
			}
	}
}
