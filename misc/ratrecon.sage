# Function for reconstruction of rational numbers from CRT data.
# Don't use this code! Use Sage's built-in rational_reconstruction().

def makeSmall(A,B):
    """
    makeSmall returns (a, b, d) such that a*A + b*B = d with b and d small.
    
    If A is the product of prime numbers and B the CRT of residues
    modulo these primes, then d/b is a rational lift of the CRT problem
    with small height.
    """
    if A < B:
        return False

    a=1; d=A
    if B == 0:
        return a, 0, d
    else:
        b1 = 0; b3 = B
        sb = b1; sd = d

    while True:
        if b3 == 0:
            #return (d-b1*B)/A, b1, d
            return (sd-sb*B)/A, sb, sd
        else:
            q = floor(d/b3)
            a, d, b1, b3 = b1, b3, a-q*b1, d.mod(b3*ZZ)
            #print("b={0}, d={1}, height(b,d)={2}\n".format(b1, d, height(b1,d)))
            if height(b1,d) <= height(sb,sd):
                sb, sd = b1, d

def height(a, b):
    return max(abs(a),abs(b))

def ratrecon(resid, moduli):
    """
    Given
        resid -- a list of residues and 
        moduli -- a list of the corresonding moduli
    returns a rational reconstruction of small height.
    """
    B = CRT(resid, moduli)
    A = mul(moduli)
    a, b, d = makeSmall(A, B)
    return d/b

## 2 examples
#moduli_91=[179,127,113,107,89,79,59,53,43]
#residues_91=[118,62,73,24,69,65,16,19,41]
#ratrecon(resid, moduli)

#moduli_52=[5,7,19,47,59,71,83,151,13,23,37,41,157]
#residues_52=[1,6,3,20,40,66,15,141,10,21,11,2,127]
#ratrecon(resid, moduli)
