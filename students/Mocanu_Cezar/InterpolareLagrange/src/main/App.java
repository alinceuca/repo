package main;

import datastructure.Polynom;
import interpolare.Lagrange;

public class App {
		
	public static void main(String args[]) throws Exception {
		
		System.out.println(Lagrange.readDataAndInterpolateForValue("input.txt", 4));
		Polynom poly = Lagrange.readDataAndGetInterpolationPolynom("input.txt");
		
		System.out.println(poly);
		System.out.println(poly.evaluate(4));
	
		
		Lagrange.readDataAndGetCoef("input.txt");
	}
}
