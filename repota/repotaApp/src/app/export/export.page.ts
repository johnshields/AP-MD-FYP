import {Component, OnInit} from '@angular/core';
import {JobReportService} from '../services/api-service';
import {ActivatedRoute, Router} from '@angular/router';
import * as jspdf from 'jspdf';
import domtoimage from 'dom-to-image';

/**
 * @author John Shields
 * @title Export Page
 * @desc Allows a user to export a report to a PDF.
 *
 * References
 * https://www.npmjs.com/package/dom-to-image
 * https://www.npmjs.com/package/jspdf
 */

@Component({
    selector: 'app-display-report',
    templateUrl: './export.page.html',
    styleUrls: ['./export.page.scss'],
})
export class ExportPage implements OnInit {
    report: any = [];
    private errorMessage;
    smallScreen = window.innerWidth <= 600;
    pdfWidth: number;
    pdfHeight: number;
    imgWidth: number;
    imgHeight: number;

    constructor(private api: JobReportService, private route: ActivatedRoute) {
    }

    /**
     * @title Error message Handlers
     * @desc Functions are used to set and get error message for error responses.
     */
    setErrorMessage(error: String) {
        this.errorMessage = error;
    }

    getErrorMessage() {
        return this.errorMessage;
    }

    // Get requested report by its ID from API.
    ngOnInit() {
        this.api.getReportById(this.route.snapshot.params['jobReportId']).subscribe(data => {
            console.log(this.route.snapshot.params['jobReportId']);
            this.report = data;
            console.log(data);
            this.setErrorMessage(''); // clear error message.
        }, error => {
            let errorMessage = JSON.stringify(error.error.messages);
            this.setErrorMessage(errorMessage);
            console.log(error);
        });

        // Resize values for mobile phones.
        if (this.smallScreen) {
            this.pdfWidth = 270;
            this.pdfHeight = 1080;
            this.imgHeight = 297;
            this.imgWidth = 100;
        } // Computer Monitors.
        else {
            this.pdfWidth = 650;
            this.pdfHeight = null;
            this.imgHeight = 297;
            this.imgWidth = 210;
        }
    }

    // Export report to a PDF.
    onExportPDF() {
        const content = document.getElementById('job-report');
        const options = {background: 'white', width: this.pdfWidth, height: this.pdfHeight, quality: 0.98};
        domtoimage.toPng(content, options).then(
            (dataUrl) => {
                // Setup PDF dimensions.
                const doc = new jspdf.jsPDF('portrait', 'mm', 'a4', true);
                doc.addImage(dataUrl, 'jpeg', 0, 0,  this.imgWidth, this.imgHeight);
                doc.save('job_report.pdf');
            });
    }
}
