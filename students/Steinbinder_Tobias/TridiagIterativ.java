package main;

import java.util.Scanner;

public class Main {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		int n;
		double[] a, b, c, d, r, s, x;
		double num;
		n = sc.nextInt();
		a = new double[n];
		b = new double[n];
		c = new double[n];
		d = new double[n];
		r = new double[n];
		s = new double[n];
		x = new double[n];
		for (int i = 0; i < n; i++) {
			a[i] = sc.nextDouble();
			b[i] = sc.nextDouble();
			c[i] = sc.nextDouble();
			d[i] = sc.nextDouble();
		}

		r[0] = -c[0] / a[0];
		s[0] = d[0] / a[0];
		for (int i = 1; i < n; i++) {
			num = (b[i] * r[i - 1] + a[i]);
			r[i] = -c[i] / num;
			s[i] = (d[i] - b[i] * s[i - 1]) / num;
		}
		x[n - 1] = s[n - 1];
		for (int i = n - 2; i >= 0; i--)
			x[i] = r[i] * x[i + 1] + s[i];

		for (int i = 0; i < n; i++)
			System.out.println("x" + i + "= " + x[i]);
		sc.close();

	}
}
