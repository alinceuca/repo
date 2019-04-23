package datastructure;

import java.text.DecimalFormat;

public class Polynom {

	private double[] coeficienti;
	private int gradMaxim;
	private static DecimalFormat df = new DecimalFormat("0.00");
	
	public Polynom(int gradMaxim, double[] coeficienti) {

		this.gradMaxim = gradMaxim;

		this.coeficienti = new double[this.gradMaxim];

		for (int i = 0; i < this.gradMaxim; i++)
			this.coeficienti[i] = coeficienti[i];
	}

	public Polynom(int gradMaxim) {
		this.gradMaxim = gradMaxim;
		coeficienti = new double[this.gradMaxim];
	}
	
	public double evaluate(double x) {
		double result = 0.0;
		
		for(int i=0;i<gradMaxim;i++)
			result += coeficienti[i] * Math.pow(x, i);
		
		return result;
	}
	
	public void scale(double alfa) {
		
		for(int i = 0;i < gradMaxim;i++)
			coeficienti[i] *= alfa;
		
	}

	public void add(Polynom q) {

		int gQ = q.getGradMaxim();
		double cQ[] = q.getCoeficienti();

		int newGradMaxim = Math.max(gradMaxim, gQ);
		double[] newCoeficienti = new double[newGradMaxim];

		for (int i = 0; i < gQ; i++)
			newCoeficienti[i] += cQ[i];

		for (int i = 0; i < gradMaxim; i++)
			newCoeficienti[i] += coeficienti[i];

		coeficienti = new double[newGradMaxim];

		for (int i = 0; i < newGradMaxim; i++)
			coeficienti[i] = newCoeficienti[i];

		gradMaxim = newGradMaxim;
	}

	public void multiply(Polynom q) {

		boolean qEsteMaiMare = false;
		
		int gQ = q.getGradMaxim();
		double cQ[] = q.getCoeficienti();

	
		
		if(gQ > gradMaxim) 
			qEsteMaiMare = true;
		
		int newGradMaxim = gradMaxim + gQ;
		double[] newCoeficienti = new double[newGradMaxim];

		if(qEsteMaiMare) 
			for (int i = 0; i < gQ; i++) //mare
				for (int j = 0; j < gradMaxim; j++)//mic
					newCoeficienti[i+j] += cQ[i] * coeficienti[j];
		else
			for (int i = 0; i < gradMaxim; i++) //mare
				for (int j = 0; j < gQ; j++)//mic
					newCoeficienti[i+j] += coeficienti[i] * cQ[j];
		

		coeficienti = new double[newGradMaxim];

		for (int i = 0; i < newGradMaxim; i++)
			coeficienti[i] = newCoeficienti[i];

		gradMaxim = newGradMaxim;
	}

	public int getGradMaxim() {
		return gradMaxim;
	}

	public double[] getCoeficienti() {
		return coeficienti;
	}
	private String sgn(double a) {
		if(a >= 0)
			return " + ";
		return " - ";
	}
	public String toString() {
		String s = "";
		
		for(int i = 0; i < gradMaxim;i++) 
			if(Math.round(coeficienti[i]) != 0)
				s += sgn(coeficienti[i]) + df.format(Math.abs(coeficienti[i])) + " X ^" + i + " "; 
		
		return s;
	}
	

}
