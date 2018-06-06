var createExpenseButton = document.getElementById('add-expense-button');
createExpenseButton.addEventListener('click', function(){
	createExpenseButton.style.display = 'none';
	document.getElementById('add-expense-form').style.display = 'block';
});