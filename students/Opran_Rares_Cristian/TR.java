package x;
import java.util.Scanner;
public class TR {

	
		int n;
		static double[] a, b, c, d, r, s, x;
		 static double num;
		 public static void main(String[] args) {
				Scanner sc = new Scanner(System.in);

				n = sc.nextInt();
				a = new double[n];
				b = new double[n];
				c = new double[n];
				d = new double[n];
				r = new double[n];
				s = new double[n];
				x = new double[n];
				
				for (int i = 0; i <= n - 1; i++) {
					a[i] = sc.nextDouble();
					b[i] = sc.nextDouble();
					c[i] = sc.nextDouble();
					d[i] = sc.nextDouble();
				}
				r[0] = -c[0] / a[0];
				s[0] = d[0] / a[0];
				recursvitate(1, true);
				for (int i = 0; i <= n - 1; i++)
					System.out.println("x" + i + "= " + x[i]);
				sc.close();
			}
		public static void recursivitate(int i, boolean f) {
			if (f)
				if (i < n - 1) {
					num = (b[i] * r[i - 1] + a[i]);
					r[i] = -c[i] / num;
					s[i] = (d[i] - b[i] * s[i - 1]) / num;
					recursivitate(i + 1, true);
				} else {
					x[n - 1] = s[n - 1];
					recursivitate(i - 1, false);
				}
			else if (i >= 0 && !f) {
				x[i] = r[i] * x[i + 1] + s[i];
				recursivitate(i - 1, false);
			}
		}

		

	

}
