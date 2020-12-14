import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AboutUsComponent } from "./about-us/about-us.component";
import { ContactUsComponent } from "./contact-us/contact-us.component";
import { FaqComponent } from "./faq/faq.component";
import { PrivacyPolicyComponent } from "./privacy-policy/privacy-policy.component";
import { HomeComponent } from "./home/home.component";
import {ShowSubsComponent} from "./show-subs/show-subs.component";
import {ArchiveUploadComponent} from "./archive-upload/archive-upload.component";
import {AdminHomeComponent} from "./admin-home/admin-home.component";

const routes: Routes = [
  { path: 'about-us', component: AboutUsComponent},
  { path: 'contact-us', component: ContactUsComponent},
  { path: 'faq', component: FaqComponent},
  { path: 'privacy-policy', component: PrivacyPolicyComponent},
  { path: 'home', component: HomeComponent},
  { path: 'admin', component: AdminHomeComponent},
  { path: 'admin/subs', component: ShowSubsComponent},
  { path: 'admin/archive/upload', component: ArchiveUploadComponent},
  { path: '', redirectTo: '/home', pathMatch: 'full'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
