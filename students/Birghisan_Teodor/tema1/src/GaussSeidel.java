
public class GaussSeidel {
	
	static double a[][] = { { 14, 3, 5 }, { -2, 24, -1 }, { -12, 12, 5 } };
	static double b[] = { 3, 2, 1 };
	static double norma = 0; 
	static final long nrMax = 1000000;
	static double x[], y[];
	static int n = b.length;
	static double epsilon = 0.0001;
	static long nr = 0;

	public static double Norma() {
		double maxim = Math.abs(x[0] - y[0]);
		for (int i = 1; i < n; i++) {
			double dif = Math.abs(x[i] - y[i]);
			if (dif > maxim)
				maxim = dif;
		}
		return maxim;
	}

	public static double Suma(int linie) {
		double suma = 0;
		for (int i = 0; i < n; i++) {
			if (linie != i)
				suma += a[linie][i] * x[i];
		}
		return suma;
	}

	public static void rezolvare() {
		x = new double[100];
		y = new double[100];
		do {
			for (int i = 0; i < n; i++) {
				y[i] = (1.00 / a[i][i]) * (b[i] - Suma(i));
			}
			norma = Norma();
			nr++;
			for (int i = 0; i < n; i++)
				x[i] = y[i];
		} while (norma > epsilon && nr < nrMax);
	}

	public static void afisare() {
		if (norma < epsilon)
			System.out.println("E buna");
		else
			System.out.println("Nu e buna");

		for (int i = 0; i < n; i++)
			System.out.println("x[" + i + "]: " + y[i]);
	}
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		System.out.println("Matricea:");
		System.out.print("A: ");
		for (int i = 0; i < a.length; i++) {
			for (int j = 0; j < a.length; j++)
				System.out.print(a[i][j] + " ");
			if (i < a.length - 1)
				System.out.print("\n   ");
		}
		System.out.print("\nB: ");
		for (int i = 0; i < b.length; i++)
			System.out.print(b[i] + " ");
		System.out.println("\n\nSolutii:");
		rezolvare();
		afisare();
	}

}
