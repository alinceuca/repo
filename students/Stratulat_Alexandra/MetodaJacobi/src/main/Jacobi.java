package main;

import java.util.Scanner;

public class Jacobi {

	static Scanner sc = new Scanner(System.in);	
	public static void main(String[] args) {
		int ok=0, n;
		double eps, norm, s;
		int nr_max = 10, nr_it=0;
		
		System.out.println("Numarul de ecuatii si necunoscute: ");
		n = sc.nextInt();
		
		double[] x = new double[nr_max];
		double[] y = new double[nr_max];
		double[] b = new double[nr_max];
		double[][] a = new double[nr_max][nr_max];
		
		System.out.println("Matricea sistemului: ");
		for(int i = 1; i <=n; i++)
			for(int j = 1; j <= n; j++) {
				System.out.println("a[" + i + " ][" + j + "]=");
				a[i][j] = sc.nextDouble();
			}
		
		System.out.println("Vectorul termenilor liberi: ");
		for(int i = 1; i <=n; i++) {
			System.out.println("b[" + i + "]=");
			b[i] = sc.nextDouble();
		}
		
		System.out.println("Aproximatia initiala a solutiei: ");
		for(int i = 1; i <= n; i++) {
			System.out.println("x[" + i + "]=");
			x[i] = sc.nextDouble();
		}
		System.out.println("Eroarea: ");
		eps = sc.nextDouble();
		norm = eps+1;
		
		do{
			
			for(int i = 1; i <= n; i++) {
				s = 0;
				for(int j = 1; j < n; j++)
					if(i != j)
						s += a[i][j] * x[j];
				y[i] = (b[i] - s) / a[i][i];
			}
			norm=0;
			for(int i = 1; i <= n; i++)
				norm += Math.abs(y[i] - x[i]);
			for(int i = 1; i <= n; i++)
				x[i] = y[i];
			nr_it++;
		}while(norm >= eps && nr_it < nr_max);
		
		System.out.println("Solutia aproximativa este");
		for(int i =1; i <= n; i++)
			System.out.println(y[i]);
		
		System.out.println("norm<eps");
		if(norm < eps)
			System.out.println("ok");
		else
			System.out.println("not ok");
	}
}
