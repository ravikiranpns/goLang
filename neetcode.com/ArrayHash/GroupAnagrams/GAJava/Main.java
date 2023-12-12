package github.com.HTTPS.goLang.neetcode.com.ArrayHash.GroupAnagrams;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Main
 */


public class Main {
    public static List<List<String>> groupAnagrams(String[] strs) {
    Map<String, List<String>> m = new HashMap<>();

    for (String s : strs) {
        char[] chars = s.toCharArray();
        Arrays.sort(chars);
        String sortedStr = new String(chars);

        if (!m.containsKey(sortedStr)) {
            m.put(sortedStr, new ArrayList<>());
        }

        m.get(sortedStr).add(s);
    }
    return new ArrayList<>(m.values());
    }
    
    public static void main(String[] str) {
        String[] strs = {"ar", "aa", "ra"};
        
       System.out.println(groupAnagrams(strs));
    }
    
}