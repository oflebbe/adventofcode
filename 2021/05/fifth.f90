program fifth
    implicit none
    integer a, b, c, d
    character*20 e
    INTEGER :: IOstatus = 0
    integer i, j, count
    integer, dimension(1000,1000),save :: arr

    interface
        subroutine handle_input(arr,x,y,z,t)
        implicit none
        integer ,intent(out) :: arr(:,:)
        integer ,intent(in) :: x,y,z,t
        end subroutine
    end interface

    arr = 0
    open(88, FILE='input.txt')
    do while (IOSTATUS .eq. 0 )
        read (88,*,IOSTAT=IOstatus)   a,b,e,c,d
        if (IOstatus .eq. 0) then
            call handle_input(arr,a,b,c,d)
        end if
    end do
    count = 0
    do  i=1,ubound(arr,1)
        !print *,arr(i,:)
        do j=1,ubound(arr,1)
            if (arr(i,j) >= 2) then
                count= count+1
            endif
        end do
    end do
    

    print *,count

end program fifth

integer function sigma(x)
    implicit none
    integer x

    if (x > 0) then
        sigma = 1
        return
    else if (x == 0) then
        sigma = 0
        return
    else
        sigma = -1
        return
    endif
end function

subroutine handle_input(arr,x,y,z,t)
    implicit none
    integer ,intent(out) :: arr(:,:)
    integer ,intent(in) :: x,y,z,t
    integer sigma
    integer  rr,wr,vr,i
    
    wr = z-x
    vr = t-y
    rr = max( abs(wr),abs(vr))

    do i=0,rr
        arr(x +i*sigma(wr)+1,y+i*sigma(vr)+1)= arr(x +i*sigma(wr)+1,y+i*sigma(vr)+1)+1
    end do

end subroutine
