package main

import (
    "encoding/json"
    "github.com/jung-kurt/gofpdf"
    "io/ioutil"
    "log"
    "os"
)

type PDFData struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

func readJSONFile(filePath string) (*PDFData, error) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    var pdfData PDFData
    err = json.Unmarshal(data, &pdfData)
    if err != nil {
        return nil, err
    }

    return &pdfData, nil
}
func createPDF(pdfData *PDFData) {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    
    // Set margins
    leftMargin, topMargin, rightMargin, bottomMargin := 10.0, 10.0, 10.0, 10.0
    pdf.SetMargins(leftMargin, topMargin, rightMargin)

	pageWidth, pageHeight := 210.0, 297.0 // A4 dimensions in mm
    pdf.Rect(leftMargin, topMargin, pageWidth-(leftMargin+rightMargin), pageHeight-(topMargin+bottomMargin), "D")
    // Set font
    pdf.SetFont("Arial", "", 10)

    // Add a logo (dummy image path, replace with actual path to logo)
    pdf.ImageOptions("./img/discord.jpg", 10, 10, 40, 0, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "")

    // Write the header
    pdf.SetFont("Arial", "B", 16)
    pdf.CellFormat(190, 10, "SMART PATHOLOGY LAB", "0", 1, "C", false, 0, "")

    // Dummy data for layout
    patientInfo := map[string]string{
        "Name":   "Jalpa S. Sharma",
        "Age":    "21 Years",
        "Sex":    "Female",
        "PID":    "555",
        // Add more patient information here
    }
    
    // Patient info section
    pdf.SetFont("Arial", "", 12)
    yOffset := 30.0 // Adjust this as needed
    for key, value := range patientInfo {
        pdf.SetX(10)
        pdf.CellFormat(40, 10, key+":", "0", 0, "", false, 0, "")
        pdf.CellFormat(0, 10, value, "0", 1, "", false, 0, "")
        yOffset += 10
    }
    
    // Table headers
    pdf.SetY(yOffset)
    pdf.SetFont("Arial", "B", 12)
    pdf.CellFormat(40, 10, "Investigation", "1", 0, "C", false, 0, "")
    pdf.CellFormat(40, 10, "Result", "1", 0, "C", false, 0, "")
    pdf.CellFormat(60, 10, "Reference Value", "1", 0, "C", false, 0, "")
    pdf.CellFormat(50, 10, "Unit", "1", 1, "C", false, 0, "")
    
    // Table content
    pdf.SetFont("Arial", "", 12)
    // Dummy data for CBC results, replace with actual data
    cbcResults := []struct {
        TestName  string
        Result    string
        RefValue  string
        Unit      string
    }{
        {"HEMOGLOBIN", "12.5", "Low 13.0 - 17.0", "g/dL"},
        // Add more CBC result lines here
    }
    for _, line := range cbcResults {
        pdf.CellFormat(40, 10, line.TestName, "1", 0, "", false, 0, "")
        pdf.CellFormat(40, 10, line.Result, "1", 0, "", false, 0, "")
        pdf.CellFormat(60, 10, line.RefValue, "1", 0, "", false, 0, "")
        pdf.CellFormat(50, 10, line.Unit, "1", 1, "", false, 0, "")
    }
    
    // Footer
    pdf.SetY(-40) // Move 40mm up from bottom
    pdf.SetFont("Arial", "I", 8)
    pdf.CellFormat(0, 10, "Thank you for choosing SMART PATHOLOGY LAB", "0", 1, "C", false, 0, "")

    // Signatures
    // Add signature images similarly to how the logo was added, adjust positions as needed

    err := pdf.OutputFileAndClose("output.pdf")
    if err != nil {
        log.Fatalf("Could not create PDF: %s", err)
    }
}


func main() {
    if len(os.Args) < 2 {
        log.Fatalf("Usage: %s <json_file_path>", os.Args[0])
    }

    jsonFilePath := os.Args[1]
    pdfData, err := readJSONFile(jsonFilePath)
    if err != nil {
        log.Fatalf("Could not read JSON file: %s", err)
    }

    createPDF(pdfData)
}


