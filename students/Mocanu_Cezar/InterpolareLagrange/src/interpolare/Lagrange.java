package interpolare;

import java.io.File;
import java.io.FileReader;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

import datastructure.Polynom;
import datastructure.Point;

public class Lagrange {

	
	

	public static void readDataAndGetCoef(String inputFile) {
		
		Scanner scanner;
		
		List<Point> points = new ArrayList<>();
		double x = 0.0, y = 0.0;
		
		try {
			scanner = new Scanner(new FileReader(new File(inputFile)));
			
			int counter = 0;
			while(scanner.hasNextDouble())
			{
				counter++;
				
				if(counter % 2 != 0)
					x = scanner.nextDouble();
				else
				{
					y = scanner.nextDouble();
					Point p = new Point(x,y);
					points.add(p);
					
				}
					
			}
				
			scanner.close();
			
		} catch (Exception e) {

			e.printStackTrace();
		}
		
		for(int i = 1;i < points.size();i++)
			System.out.println("[y0,...," + "y"+(points.size()-i)+"] = "+solveNewtonCoeficients(0,points.size()-i,points));
		
	}
	
	//
	public static double solveNewtonCoeficients(int p,int u,List<Point> points) {
		
		/// punctul P
		Point p0 = points.get(p);
		
		/// punctul U
		Point p1 = points.get(u);
		if(u - p == 1) 
			return (p1.getY() - p0.getY()) / (p1.getX() - p0.getX());
		else 	
			return (solveNewtonCoeficients(p, u-1, points) - solveNewtonCoeficients(p+1, u, points)) / (p1.getX() - p0.getX());
		
	}
	
	/**
	 * @javadoc 
	 * Reads from inputFile every pair of values as points
	 * </br>
	 * x0 y0</br>
	 * ....</br>
	 * xn yn</br>
	 * 
	 * and solves the Lagrange interpolation value for Z
	 * 
	 * @param
	 * inputFile - file containing numbers under the structure
	 *</br>
	 * x0 y0</br>
	 * ....</br>
	 * xn yn</br>
	 * 
	 * @param
	 * Z - number to interpolate for
	 * 
	 */
	public static double readDataAndInterpolateForValue(String inputFile, double Z) {

		Scanner scanner;
		
		List<Point> points = new ArrayList<>();
		double x = 0.0, y = 0.0;
		
		try {
			scanner = new Scanner(new FileReader(new File(inputFile)));
			
			int counter = 0;
			while(scanner.hasNextDouble())
			{
				counter++;
				
				if(counter % 2 != 0)
					x = scanner.nextDouble();
				else
				{
					y = scanner.nextDouble();
					Point p = new Point(x,y);
					points.add(p);
					
				}
					
			}
				
			scanner.close();
			
		} catch (Exception e) {

			e.printStackTrace();
		}
		
		return solveIntepolationForValue(points,Z);
		

	}

	private static double solveIntepolationForValue(List<Point> points,double Z) {
		
		double out = 0.0;
		int numberOfPoints = points.size();
		
		if(numberOfPoints == 0)
			return Integer.MIN_VALUE;
		
		for(int i=0;i<numberOfPoints;i++) {
		
			double productNumerator = 1;
			double productDenominator = 1;
			double xi = points.get(i).getX();
			
			for(int j=0;j<numberOfPoints;j++) {
				if(i == j)
					continue;
				productNumerator *= (Z - points.get(j).getX());
				productDenominator *= (xi - points.get(j).getX());
			}
			
			out += points.get(i).getY() * (productNumerator/productDenominator);
		}
		
		return out;
	}

	
	/**
	 * @javadoc 
	 * Reads from inputFile every pair of values as points
	 * and generates an aproximation of Lagrange Interpolation Polynom
	 * 
	 * @param
	 * inputFile - file containing numbers under the structure
	 *</br>
	 * x0 y0</br>
	 * ....</br>
	 * xn yn</br>
	 * 
	 * 
	 */
	public static Polynom readDataAndGetInterpolationPolynom(String inputFile) {

		Scanner scanner;
		
		List<Point> points = new ArrayList<>();
		double x = 0.0, y = 0.0;
		
		try {
			scanner = new Scanner(new FileReader(new File(inputFile)));
			
			int counter = 0;
			while(scanner.hasNextDouble())
			{
				counter++;
				
				if(counter % 2 != 0)
					x = scanner.nextDouble();
				else
				{
					y = scanner.nextDouble();
					Point p = new Point(x,y);
					points.add(p);
					
				}
					
			}
				
			scanner.close();
			
		} catch (Exception e) {

			e.printStackTrace();
		}
		
		return getInterpolationPolynom(points);
		

	}
	

	private static Polynom getInterpolationPolynom(List<Point> points) {

		Polynom lagrangeInterpolationPolynom = new Polynom(0);
		int numberOfPoints = points.size();
		
		if(numberOfPoints == 0)
			return new Polynom(0);
		
		for (int i = 0; i < numberOfPoints; i++) {

			double scalar = 0.0;
			double numitor = 1.0;
			double xi = points.get(i).getX();
			Polynom poly = null;

			for (int j = 0; j < numberOfPoints; j++) {
				if (i == j)
					continue;
				if (poly == null) {
					double coef[] = { -points.get(j).getX(), 1 };
					poly = new Polynom(coef.length, coef);
				} else {

					double coef[] = { -points.get(j).getX(), 1 };
					Polynom q = new Polynom(coef.length, coef);
					poly.multiply(q);

				}

				numitor *= xi - points.get(j).getX();
			}

			scalar = points.get(i).getY() / numitor;

			poly.scale(scalar);

			lagrangeInterpolationPolynom.add(poly);

		}

		return lagrangeInterpolationPolynom;

	}
	
	

}
