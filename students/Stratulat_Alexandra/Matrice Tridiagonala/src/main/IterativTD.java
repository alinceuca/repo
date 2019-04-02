package main;

import java.util.Scanner;

public class IterativTD {

	static Scanner sc = new Scanner(System.in);
	
	public static void main(String[] args) {
		double[] a, b, c, d, x, r, s;
		
		System.out.println("n=");
		int n = sc.nextInt();
		
		a = new double[n];
		b = new double[n];
		c = new double[n];
		d = new double[n];
		x = new double[n];
		r = new double[n];
		s = new double[n];
		
		for(int i = 0; i < n; i++) {
			System.out.println("a[" + i + "]=");
			a[i] = sc.nextDouble();
		}
		
		for(int i = 1; i < n; i++) {
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
		
		
		r[0] = -c[0]/a[0];
		s[0] = d[0]/a[0];
		
		for(int i = 1; i <= n-2; i++) {
			r[i] = -c[i]/(b[i]*r[i-1]+a[i]);
			s[i] = (d[i]-b[i]*s[i-1])/(b[i]*r[i-1]+a[i]);
		}
		x[n-1] = (d[n-1]-b[n-2]*s[n-2])/(b[n-2]*r[n-2]+a[n-1]);
		for(int i = n-2; i >= 0; i--)
			x[i] = r[i]*x[i+1]+s[i];
			
		for(int i = 0; i < n; i++)
			System.out.println(x[i]);
		
	}
}
