describe('historyTest', function(){
    this.beforeEach(()=>{
        cy.visit('http://localhost:8000/user/login')
        cy.get('.ant-pro-form-login-main').get('#email').type('11@111.com')
        cy.get('.ant-pro-form-login-main').get('#password').type('123456')
        cy.get('button:contains("Login")').click()
    })
    // it('proceedTest', function(){
    // })
})