import {Component, OnInit} from '@angular/core';
import {JobReport, JobReportService} from '../services/api-service';
import {NgForm} from '@angular/forms';
import {Router} from '@angular/router';

/**
 * @author John Shields
 * @title Create Page
 * @desc Allows a User to create a new report using the JobReport Model.
 */

@Component({
    selector: 'app-create',
    templateUrl: './create.page.html',
    styleUrls: ['./create.page.scss'],
})
export class CreatePage implements OnInit {
    list1: any[];
    list2: any[];
    list3: any[];
    checkBoxValue1: number;
    checkBoxValue2: number;
    checkBoxValue3: number;
    vehicles: any = [];
    private errorMessage;

    constructor(private api: JobReportService, private router: Router) {
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

    /**
     * @title Create Report
     * @desc Uses the JobReport Model to take in the input and create a report.
     */
    createReport(form: NgForm) {
        // Make the true/false values of check boxes to 1s and 0s.
        // warranty
        if (form.value.warranty === true) {
            this.checkBoxValue1 = 1;
        } else {
            this.checkBoxValue1 = 0;
        }
        // breakdown
        if (form.value.breakdown === true) {
            this.checkBoxValue2 = 1;
        } else {
            this.checkBoxValue2 = 0;
        }
        // job complete
        if (form.value.jobComplete === true) {
            this.checkBoxValue3 = 1;
        } else {
            this.checkBoxValue3 = 0;
        }

        // Use JobReport Model
        const object: JobReport = {
            date: form.value.date,
            vehicleModel: form.value.vehicleModel,
            vehicleReg: form.value.vehicleReg,
            milesOnVehicle: form.value.milesOnVehicle,
            vehicleLocation: form.value.vehicleLocation,
            warranty: this.checkBoxValue1,
            breakdown: this.checkBoxValue2,
            customerName: form.value.customerName,
            complaint: form.value.complaint,
            cause: form.value.cause,
            correction: form.value.correction,
            parts: form.value.parts,
            workHours: form.value.workHours,
            jobComplete: this.checkBoxValue3
        };

        // Push data to API to create report using the model.
        this.api.createReport(object).subscribe(data => {
            this.api.createReport(data);
            this.setErrorMessage(''); // clear error message.
            this.router.navigate(['/history']);
        }, error => {
            // Get error from response.
            let errorMessage = JSON.stringify(error.error.messages);
            this.setErrorMessage(JSON.parse(errorMessage));
            console.log(error);
        });
    }

    /**
     * @title ngOnInit
     * @desc Gets vehicle data from API & lists for form check boxes.
     */
    ngOnInit() {
        // Get vehicle data from API.
        this.api.getCarApiData().subscribe(data => {
            // Strip out the keys from data.
            for (let key in data) {
                this.vehicles = data[key];
            }
            this.setErrorMessage('');
        }, error => {
            let errorMessage = JSON.stringify(error.error.messages);
            this.setErrorMessage(errorMessage);
            console.log(error);
        });

        // Lists for check box values.
        this.list1 = [{
            id: 1,
            title: 'Warranty',
            checked: false,
            value: 0
        }];
        this.list2 = [{
            id: 1,
            title: 'Breakdown',
            checked: false,
            value: 0
        }];
        this.list3 = [{
            id: 1,
            title: 'Complete',
            checked: false,
            value: 0
        }];
    }
}
