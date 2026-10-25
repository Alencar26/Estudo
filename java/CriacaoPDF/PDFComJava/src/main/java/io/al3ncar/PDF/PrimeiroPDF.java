package io.al3ncar.PDF;

import com.lowagie.text.Document;
import com.lowagie.text.Paragraph;
import com.lowagie.text.pdf.PdfWriter;

import java.io.FileOutputStream;

public class PrimeiroPDF {

    public PrimeiroPDF(String texto) {
        Document documentoPDF = new Document();

        try {
            PdfWriter.getInstance(documentoPDF, new FileOutputStream("relatorio1.pdf"));
            documentoPDF.open();
            Paragraph paragrafoTexte = new Paragraph(texto);
            documentoPDF.add(paragrafoTexte);
        } catch (Exception e) {
            System.out.println(e.getMessage());
        } finally {
            documentoPDF.close();
        }
    }
}
