package Gauss;

public class Gauss {

	static double A[][] = { { 28, 1, 4 }, { 11, 23, -8 }, { 26, 27, 54 } };
	static double norm = 0;
	static double b[] = { 21, 23, 24 };
	static int n = b.length;
	static double eps = 0.0001;
	static long nr_it = 0;
	static final long nr_max = 1000000;
	static double x[], y[];

	public static double calcNorm() {
		double max = Math.abs(x[0] - y[0]);
		for (int i = 1; i < n; i++) {
			double dif = Math.abs(x[i] - y[i]);
			if (dif > max)
				max = dif;

		}
		return max;
	}

	public static double calcSum1(int linie) {
		double sum = 0;
		for (int i = 0; i < linie; i++) {
			if (linie != i)
				sum += A[linie][i] * y[i];
		}
		return sum;
	}

	public static double calcSum2(int linie) {
		double sum = 0;
		for (int i = linie; i < n; i++) {
			if (linie != i)
				sum += A[linie][i] * x[i];
		}
		return sum;
	}

	public static void solve() {
		x = new double[100];
		y = new double[100];
		do {
			for (int i = 0; i < n; i++) {
				y[i] = (1.00 / A[i][i]) * (b[i] - calcSum1(i) - calcSum2(i));
			}
			norm = calcNorm();
			nr_it++;
			for (int i = 0; i < n; i++)
				x[i] = y[i];
		} while (norm > eps && nr_it < nr_max);
	}

	public static void afis() {
		if (norm < eps)
			System.out.println("OK");
		else
			System.out.println("!OK");

		for (int i = 0; i < n; i++)
			System.out.println("x[" + i + "]: " + y[i]);
	}

	public static void main(String args[]) {
		System.out.println("Matricea:");
		System.out.print("A: ");
		for (int i = 0; i < A.length; i++) {
			for (int j = 0; j < A.length; j++)
				System.out.print(A[i][j] + " ");
			if (i < A.length - 1)
				System.out.print("\n   ");
		}
		System.out.print("\nB: ");
		for (int i = 0; i < b.length; i++)
			System.out.print(b[i] + " ");
		System.out.println("\n\nSolutii:");
		solve();
		afis();
	}
}