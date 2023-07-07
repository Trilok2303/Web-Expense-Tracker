function addExpense() {
    var descriptionInput = document.getElementById('expense-description');
    var amountInput = document.getElementById('expense-amount');
    var expensesList = document.getElementById('expenses');

    var description = descriptionInput.value;
    var amount = amountInput.value;

    if (description === '') {
        alert('Please enter a description for the expense.');

        return;
    }
    if (amount === '') {
        alert('Please enter amount for the expense.');

        return;
    }
    var expenseItem = document.createElement('li');
    var expenseText = document.createElement('span');
    expenseText.textContent = description;
    expenseItem.appendChild(expenseText);

    var expenseAmount = document.createElement('span');
    expenseAmount.textContent = '$' + amount;
    expenseItem.appendChild(expenseAmount);

    expensesList.appendChild(expenseItem);

    descriptionInput.value = '';
    amountInput.value = '';
}