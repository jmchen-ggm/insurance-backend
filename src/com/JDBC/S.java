package com.JDBC;
import java.sql.*;

public class S
{
    public static void main( String args[] )
    {
        Connection c = null;
        Statement stmt = null;
        try {
            Class.forName("org.sqlite.JDBC");
            c = DriverManager.getConnection("jdbc:sqlite:insuranceDB.db");
            c.setAutoCommit(false);
            System.out.println("Opened database successfully");

            stmt = c.createStatement();
            // String sql = "DELETE from produce where rowid=2;";
            //  stmt.executeUpdate(sql);
            //  c.commit();

            ResultSet rs = stmt.executeQuery( "SELECT * FROM produce;" );
            while ( rs.next() ) {
                //int id = rs.getInt("rowid");
                String  name = rs.getString("name");
                int age  = rs.getInt("rank");
                int  address = rs.getInt("money");
                float salary = rs.getFloat("score");
                //System.out.println( "rowid = " + id );
                System.out.println( "name = " + name );
                System.out.println( "rank = " + age );
                System.out.println( "money = " + address );
                System.out.println( "score = " + salary );
                System.out.println();
            }
            rs.close();
            stmt.close();
            c.close();
        } catch ( Exception e ) {
            System.err.println( e.getClass().getName() + ": " + e.getMessage() );
            System.exit(0);
        }
        System.out.println("Operation done successfully");
    }
}