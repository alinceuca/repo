package main;

import java.util.Scanner;

public class RecursivTD {
	static Scanner sc = new Scanner(System.in);
	static double[] a, b, c, d, x, r, s;
	static int n;

	public static void recRS(int i) {
		if (i == 0 && a[0] != 0) {
			r[i] = -c[i] / a[i];
			s[i] = d[i] / a[i];
		}
		if (i < n-1) {
			r[i] = -c[i]/(b[i]*r[i-1]+a[i]);
			s[i] = (d[i]-b[i]*s[i-1])/(b[i]*r[i-1]+a[i]);
			recRS(i+1);
		}
	}

	public static void recXA(int i) {
		if (i == n-1) {
			x[i] = (d[i]-b[i-1]*s[i-1])/(b[i-1]*r[i-1]+a[i]);
			recXA(i-1);
		} else if (i >= 0) {
			x[i] = r[i]*x[i+1]+s[i];
			recXA(i-1);
		}
	}
	
	public static void recursiv(int i) {
		if (i == 0 && a[0] != 0) {
			r[i] = -c[i] / a[i];
			s[i] = d[i] / a[i];
		}
		if (i < n-1) {
			r[i] = -c[i]/(b[i]*r[i-1]+a[i]);
			s[i] = (d[i]-b[i]*s[i-1])/(b[i]*r[i-1]+a[i]);
			recursiv(i+1);
		}
		else {
			x[n] = (d[n]-b[n-1]*s[n-1])/(b[n-1]*r[n-1]+a[n]);
			recursiv(i-1);
		}
		if(i >= 0) {
			x[i] = r[i]*x[i+1]+s[i];
			recursiv(i-1);
		}
	}

	

	

	public static void main(String[] args){

		Scanner sc = new Scanner(System.in);
		System.out.println("n=");
		n = sc.nextInt();
		a = new double[n];
		b = new double[n];
		c = new double[n];
		d = new double[n];
		x = new double[n+1];
		r = new double[n];
		s = new double[n];
		
		for(int i = 0; i < n; i++) {
			System.out.println("a[" + i + "]=");
			a[i] = sc.nextDouble();
		}
		
		for(int i = 1; i<n; i++) {
			System.out.println("b[" + i + "]=");
			b[i] = sc.nextDouble();
		}
		
		for(int i = 0; i < n-1; i++) {
			System.out.println("c[" + i + "]=");
			c[i] = sc.nextDouble();
		}
		
		for(int i = 0; i <n; i++) {
			System.out.println("d[" + i + "]=");
			d[i] = sc.nextDouble();
		}

		
		recRS(0);
		recXA(n-1);
		
		System.out.println("\nSolutii:");
		for (int i = 0; i < n; i++)
			System.out.println("x[" + i + "]: " + x[i]);
	}
}
