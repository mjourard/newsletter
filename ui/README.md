# Ui

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 10.0.6.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory. Use the `--prod` flag for a production build.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via [Protractor](http://www.protractortest.org/).

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI README](https://github.com/angular/angular-cli/blob/master/README.md).

## Icons
* Get icons from [fontawesome.com](https://fontawesome.com/icons)
  * select from *free*,*solid*,*regular* and *brands
  * remove the **fa-** from the start when adding it to the `[icons]` directive

## Notes
#### 2020-10-12
* enabled console.log for development by changing **true** to **false** in the **no-console** section of *tslint.json*

#### 2020-10-16
* Add NG_DEPLOY_AWS_ACCESS_KEY_ID and NG_DEPLOY_AWS_SECRET_ACCESS_KEY to your environment variables for deployment

#### 2020-10-20
* Added an env service that will load in an **env.js** file from the `src/` folder, as well as an **env.js.dist** file 
that can be copied to **env.js** and prod will use those values. Values here will be different for each person deploying/environment  
