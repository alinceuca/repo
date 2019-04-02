
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.Scanner;

public class Jacobi {

	public static double[][]A;
	public static double b[];
	public static double X[],Y[];
	public static double nr_max=10000;
	public static int nr_it=0;
	public static double norm;
	public static double eps=Math.pow(10, -4);
	public static int n,ok;
	public static Scanner sc;
	
	static double sum(int linie)
	{
		double s=0;
		for(int i=0;i<n;i++)
				if(linie!=i)
					s+=(A[linie][i]*X[i]);
		return s;
					
	}
	
	static double norma_max()
	{
		double dif=0;
		double max=Math.abs(X[0]-Y[0]);
		for(int i=1;i<n;i++) {
			dif=Math.abs(X[i]-Y[i]);
			if(dif>max)
				max=dif;
		}
		return max;
	}
	
	static void change() {
		for(int i=0;i<n;i++)
			X[i]=Y[i];
	}
	
	static void jacobi()
	{
		do {
			for(int i=0;i<n;i++)
				Y[i]=(b[i]-sum(i))/A[i][i];
			norm=norma_max();
			nr_it++;
			change();
		}while(norm>=eps && nr_it<nr_max);
	}
	
	
	static void read() throws IOException
	{
		sc = new Scanner(new FileReader(new File("date1.in")));
		n=sc.nextInt();
		A=new double[n][n];
		b=new double[n];
		X=new double[n];
		Y=new double[n];
		for(int i=0;i<n;i++)
			for(int j=0;j<n;j++)
				A[i][j]=sc.nextDouble();
		for(int i=0;i<n;i++)
			b[i]=sc.nextDouble();
	}
	
	public static void main(String[] args) {
		try {
			read();
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		jacobi();
		for(int i=0;i<n;i++)
			System.out.println(Y[i]+ " ");
		if(norm<eps)
			ok=1;
		System.out.println("Ok: "+ ok);
		
			

	}

}
