import java.util.Scanner;

public class Main {

	static double mat[][];
	static double b[], x[], y[];
	static int n;
	static double eps;
	static long nrMax;

	static double norm() {
		double max = Math.abs(x[0] - y[0]);
		for (int i = 1; i < n; i++) {
			double calc = Math.abs(x[i] - y[i]);
			if (calc > max)
				max = calc;
		}
		return max;
	}

	static double suma(int row) {
		double suma = 0.0;
		for (int i = 0; i < n; i++)
			if (i != row)
				suma += mat[row][i] * x[i];
		return suma;
	}

	static double sumaM(int row) {
		double suma = 0.0;
		for (int i = row; i < n; i++)
			if (i != row)
				suma += mat[row][i] * x[i];
		return suma;
	}

	static double sumaU(int row) {
		double suma = 0.0;
		for (int i = 0; i < row; i++)
			if (i != row)
				suma += mat[row][i] * y[i];
		return suma;
	}

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		eps = 0.0001;
		nrMax = 10000;
		n = sc.nextInt();
		mat = new double[n][n];
		b = new double[n];
		x = new double[n];
		y = new double[n];
		double norm;
		for (int i = 0; i < n; i++)
			for (int j = 0; j < n; j++)
				mat[i][j] = Double.parseDouble(sc.next());
		for (int i = 0; i < n; i++)
			b[i] = Double.parseDouble(sc.next());
		sc.close();
		int nrCrt = 0;
		do {
			for (int i = 0; i < n; i++)
				y[i] = (1.00 / mat[i][i]) * (b[i] - sumaU(i) - sumaM(i));
			norm = norm();
			for (int i = 0; i < n; i++)
				x[i] = y[i];
			nrCrt++;
		} while (norm >= eps && nrCrt < nrMax);
		for (int i = 0; i < n; i++)
			System.out.println(y[i]);
	}
}