import java.util.Scanner;
public class Jacobi {

public static double max(double a, double b) {
	if (a > b)
		return a;
	return b;
	}

	public static double norm(double[] x, double[] y, int n) {
		double maxi = 0;
		for (int i = 0; i < n; i++) {
			maxi = max(maxi, Math.abs(x[i] - y[i]));
		}
		return maxi;
	}

public static void main(String[] args) {
	Scanner sc = new Scanner(System.in);
	System.out.println("Lungimea= ");
	int n = sc.nextInt();
	double a[][] = new double[n][n];
	int k = 1, kmax = 500;
	System.out.println("Elementele matricei: ");
	for (int i = 0; i < n; i++)
	   for (int j = 0; j < n; j++) {
		  a[i][j] = sc.nextDouble();
			}
  double y[] = new double[n];
  double x[] = new double[n];
  for (int i = 0; i < n; i++)
		x[i] = 0;
  double b[] = new double[n];
	System.out.println("b=");
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
					}
				}
		y[i] = (1 / a[i][i]) * (b[i] - c);
			}
        		k++;

	nor = norm(x, y, n);
	  for (int i = 0; i < n; i++) {
			x[i] = y[i];
			}

	} while ((nor >= ep) && (k < kmax));

		System.out.println("Solutia:");
		for (int i = 0; i < n; i++)
			System.out.println(x[i] + " ");
	}

}