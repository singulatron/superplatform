# page component

## Usage

## One column - only icons

```html
<app-page
	title="Title"
	[menuWidth]="'3rem'" <!-- Default with is 80% -->
	[columnWidths]="['3rem']"
	[mobileColumnWidths]="['3rem']"
>
	<ng-template #columnContent>
		<div>
			<app-icon-menu></app-icon-menu>
		</div>
	</ng-template>
	<!-- Add more columns like this <ng-template #columnContent> Another column </ng-template> -->
	<ng-template #mainContent> Main content here </ng-template>
</app-page>
```

### Two columns - icons + content

```html
<app-page
	title="Title"
	[menuWidth]="'90%'" <!-- Default with is 80% -->
	[columnWidths]="['3rem', '25%']"
	[mobileColumnWidths]="['3rem', '100%']"
>
	<ng-template #columnContent>
		<div>
			<app-icon-menu></app-icon-menu>
		</div>
	</ng-template>
		<!-- Add more columns like this <ng-template #columnContent> Another column </ng-template> -->
	<ng-template #mainContent> Main content here </ng-template>
</app-page>
```
