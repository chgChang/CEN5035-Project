describe('register', function(){
    this.beforeEach(()=>{
        cy.visit('http://localhost:8000/user/login')
        cy.get('.ant-tabs-tab:contains("Register")').click()
    })
    it('newTest', function(){
        cy.get('.ant-pro-form-login-main').get('#email').type('1212345@111.com')
        cy.get('.ant-pro-form-login-main').get('#username').type('cc')
        cy.get('.ant-pro-form-login-main').get('#registerPassword').type('zc1234')
        cy.get('.ant-pro-form-login-main').get('#confirmPassword').type('zc1234')
        cy.get('button:contains("Login")').click()
        cy.url().should('eq', 'http://localhost:8000/welcome/Items')
    })

    it('existTest', function(){
        cy.get('.ant-pro-form-login-main').get('#email').type('cz@cz.com')
        cy.get('.ant-pro-form-login-main').get('#username').type('cc')
        cy.get('.ant-pro-form-login-main').get('#registerPassword').type('zc1234')
        cy.get('.ant-pro-form-login-main').get('#confirmPassword').type('zc1234')
        cy.get('button:contains("Login")').click()
        cy.url().should('include', '/user/login')
    })
})