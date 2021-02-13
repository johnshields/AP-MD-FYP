import {Component, OnInit} from '@angular/core';
import {AccountService, JobReportService} from '../services/client_stubs';
import {ActivatedRoute, Router} from '@angular/router';

@Component({
    selector: 'app-report-history',
    templateUrl: './report-history.page.html',
    styleUrls: ['./report-history.page.scss'],
})
export class ReportHistoryPage implements OnInit {
    reports: any = [];
    public errorMsg: string;
    public successMsg: string;

    constructor(private api: JobReportService, private router: Router) {
    }

    ngOnInit() {
        // set dark theme to default theme
        document.body.setAttribute('color-theme', 'dark');

        // get all worker's reports
        this.api.getReports().subscribe(data => {
            this.reports = data;
            console.log('[INFO] Reports have been processed.');
            console.log(this.reports);
        });
    }

    // refresh page to see new reports
    refreshPage() {
        window.location.reload();
    }
}
