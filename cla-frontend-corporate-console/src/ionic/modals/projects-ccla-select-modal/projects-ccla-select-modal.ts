// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT

import { Component, ChangeDetectorRef } from "@angular/core";
import {
  NavController,
  NavParams,
  ModalController,
  ViewController,
  AlertController,
  IonicPage
} from "ionic-angular";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { ClaService } from "../../services/cla.service";
import { ClaCompanyModel } from "../../models/cla-company";

@IonicPage({
  segment: "projects-ccla-select-modal"
})
@Component({
  selector: "projects-ccla-select-modal",
  templateUrl: "projects-ccla-select-modal.html"
})
export class ProjectsCclaSelectModal {
  projects: any;
  loading: any;
  company: ClaCompanyModel;

  constructor(
    public navParams: NavParams,
    public navCtrl: NavController,
    public viewCtrl: ViewController,
    public formBuilder: FormBuilder,
    private claService: ClaService
  ) {
    this.getDefaults();
  }

  getDefaults() {
    this.loading = true;
    this.company = this.navParams.get("company");
  }

  ngOnInit() {
    this.getProjectsCcla();
  }

  getProjectsCcla() {
    this.claService.getCompanyUnsignedProjects(this.company.company_id).subscribe(response => {
      this.loading = false;
      this.projects = response;
    });

  }

  selectProject(project) {
    this.navCtrl.push("AuthorityYesnoPage", {
      projectId: project.project_id,
      company: this.company
    });

    this.dismiss();
  }

  dismiss() {
    this.viewCtrl.dismiss();
  }
}
