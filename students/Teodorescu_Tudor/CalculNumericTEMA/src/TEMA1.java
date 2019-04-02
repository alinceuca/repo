import java.util.Scanner;
public class TEMA1 {
	
	public static double suma(double a[],double x[],int n,int j,int i)
	{
		double s=0;
		for(j=1;j<=n;j++)
		{
			if(j!=i)
				s=s+a[j]*x[j];
			
		}
		return s;
	}
	public static double max(double a,double b)
	{
		if(a>b)
			return a;
		return b;
		
	}
	public static double maxsir(double x[],double y[],int n,int i)
	{
		if(i==n)
			return x[i]-y[i];
		else
			return max((x[i]-y[i]),maxsir(x,y,n,i+1));
	}
	public static void main(String[] args) {
		
		Scanner sc = new Scanner(System.in);
		System.out.print("n=");
		int n = sc.nextInt();
		int nrmax = n;
		int j=1,i=0,nr_it=0,ok=0,eps = 1/(10*10*10);
		double norm;
		double a[] = new double [100];
		double y[] = new double [100];
		double b[] = new double[100];
		double x[] = new double [100];
		for(int l=1;l<=n;l++)
			a[l] = sc.nextInt();
		for(int l=1;l<=n;l++)
			x[l] = sc.nextInt();
		for(int l=1;l<=n;l++)
			b[l] = sc.nextInt(); 
		do {
			if(i<=n)
				i++;
			y[i] = (b[i]-suma(a,x,n,j,i))/a[i];
			norm = maxsir(x,y,n,i);
			nr_it++;
			x[i]=y[i];
		}while(norm>=eps&&nr_it<nrmax);
		
		for(int l=1;l<=n;l++)
			System.out.print(y[l]+" ");
		System.out.println();
		if(norm<eps)
			ok=1;
		if(ok==1)
			System.out.print("Norma e mai mica decat epsilon");
		else
			System.out.print("S-a iesit din 'do-while' prin depasirea nr maxim");
		
		
		
		
	}

}
