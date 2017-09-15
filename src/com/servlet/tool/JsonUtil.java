package com.servlet.tool;

import com.google.gson.JsonArray;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParser;
import com.beans.produce;

import javax.servlet.http.HttpServletRequest;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.sql.Date;
import java.text.Format;
import java.text.SimpleDateFormat;
import java.util.List;

public class JsonUtil {

    public static JsonObject String2Json(String data) {
        JsonParser parser = new JsonParser();
        return (JsonObject) parser.parse(data);
    }

    public static String getPostBody(HttpServletRequest request) throws IOException {
        // get the body of the request object to a string

        String body = null;
        StringBuilder stringBuilder = new StringBuilder();
        BufferedReader bufferedReader = null;

        try {
            InputStream inputStream = request.getInputStream();
            if (inputStream != null) {
                bufferedReader = new BufferedReader(new InputStreamReader(inputStream));
                char[] charBuffer = new char[128];
                int bytesRead = -1;
                while ((bytesRead = bufferedReader.read(charBuffer)) > 0) {
                    stringBuilder.append(charBuffer, 0, bytesRead);
                }
            } else {
                stringBuilder.append("");
            }
        } catch (IOException ex) {
            ex.printStackTrace();
            throw ex;
        } finally {
            if (bufferedReader != null) {
                try {
                    bufferedReader.close();
                } catch (IOException ex) {
                    throw ex;
                }
            }
        }

        body = stringBuilder.toString();
        return body;
    }

    public static String retDefaultJson(boolean state, String errMsg, String out_put, JsonElement element){
        JsonObject object = new JsonObject();

        object.addProperty("state",state);
        object.addProperty("errMsg", errMsg);
        object.addProperty("out_put", out_put);
        object.add("element", element);

        return object.toString();
    }

    public static String toDateFormat(Date date){
        Format format = new SimpleDateFormat("yyyy-MM-dd hh:mm:ss");
        return format.format(date);
    }


    public static JsonObject Produce2Json(produce Produce){
        JsonObject object = new JsonObject();
        object.addProperty("id",Produce.getId());
        object.addProperty("name",Produce.getName()== null ? "No Name" : Produce.getName());
        object.addProperty("addition",Produce.getAddition() == null ? "No Addition" : Produce.getAddition());
        object.addProperty("money",Produce.getMoney());
        object.addProperty("rank",Produce.getRank());
        object.addProperty("score",Produce.getScore());
        object.addProperty("lgroup",Produce.getLgroup());
        object.addProperty("lkind",Produce.getLkind());
        object.addProperty("lnumber",Produce.getLnumber());
        object.addProperty("lpercenta",Produce.getLpercenta());
        object.addProperty("lpercentb",Produce.getLpercentb());
        object.addProperty("sgroup",Produce.getSgroup());
        object.addProperty("skind",Produce.getSkind());
        object.addProperty("snumber",Produce.getSnumber());

        return object;
    }

}
