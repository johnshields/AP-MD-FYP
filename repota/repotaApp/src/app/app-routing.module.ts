import {NgModule} from '@angular/core';
import {PreloadAllModules, RouterModule, Routes} from '@angular/router';
import {AuthGuard} from './services/auth-service/auth.guard';

const routes: Routes = [
    {
        path: '',
        loadChildren: () => import('./account/account.module').then(m => m.AccountPageModule)
    },
    {
        path: 'about',
        loadChildren: () => import('./about/about.module').then(m => m.AboutPageModule)
    },
    {
        path: 'home',
        loadChildren: () => import('./home/home.module').then(m => m.HomePageModule),
        canActivate: [AuthGuard] // Block User if they are not logged in (no cookie).
    },
    {
        path: 'create',
        loadChildren: () => import('./create/create.module').then(m => m.CreatePageModule),
        canActivate: [AuthGuard]
    },
    {
        path: 'history',
        loadChildren: () => import('./history/history.module').then(m => m.HistoryPageModule),
        canActivate: [AuthGuard]
    },
    {
        path: 'display',
        loadChildren: () => import('./display/display.module').then(m => m.DisplayPageModule),
        canActivate: [AuthGuard]
    },
    {
        path: 'display/:jobReportId',
        loadChildren: './display/display.module#DisplayPageModule',
        canActivate: [AuthGuard]
    },
    {
        path: 'export',
        loadChildren: () => import('./export/export.module').then(m => m.ExportPageModule),
        canActivate: [AuthGuard]
    },
    {
        path: 'export/:jobReportId',
        loadChildren: './export/export.module#ExportPageModule',
        canActivate: [AuthGuard]
    },
    {
        path: 'login',
        loadChildren: () => import('./login/login.module').then(m => m.AccountPageModule)
    },
    {
        path: 'register',
        loadChildren: () => import('./register/register.module').then(m => m.RegisterPageModule)
    },
    {
        path: 'edit',
        loadChildren: () => import('./edit/edit.module').then(m => m.EditPageModule),
        canActivate: [AuthGuard]
    },
    {
        path: 'edit/:jobReportId',
        loadChildren: './edit/edit.module#EditPageModule',
        canActivate: [AuthGuard]
    }
];

@NgModule({
    imports: [
        RouterModule.forRoot(routes, {preloadingStrategy: PreloadAllModules})
    ],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
