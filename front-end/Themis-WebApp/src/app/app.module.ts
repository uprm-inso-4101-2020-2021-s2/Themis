import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './components/Home/home/home.component';
import { GroupsComponent } from './components/User-Account/groups/groups.component';
import { ManageComponent } from './components/User-Account/groups/Groups-components/Manage/manage/manage.component';
import { GypoComponent } from './components/User-Account/groups/Groups-components/GroupsYourPartOf/gypo/gypo.component';
import { SearchBarComponent } from './components/User-Account/groups/Groups-components/Search/search-bar/search-bar.component';
import { HowToComponent } from './components/Home/home/HowTo/how-to/how-to.component';
import { InformationComponent } from './components/Home/home/Information/information/information.component';
import { ContactUsComponent } from './components/ContactUs/contact-us/contact-us.component';
import { NavbarComponent } from './components/Navbar/navbar/navbar.component';
import { WalletComponent } from './components/Navbar/navbar/Wallet/wallet/wallet.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    GroupsComponent,
    ManageComponent,
    GypoComponent,
    SearchBarComponent,
    HowToComponent,
    InformationComponent,
    ContactUsComponent,
    NavbarComponent,
    WalletComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
